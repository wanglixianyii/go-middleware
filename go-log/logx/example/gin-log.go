package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-log/logx"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func main() {

	r := gin.New()

	lc := logx.LogConfig{
		LogLevel:          "debug",
		LogFileName:       fmt.Sprintf("%v.log", time.Now().Unix()),
		LogFileMaxSize:    1,
		LogFileMaxBackups: 5,
		LogMaxAge:         30,
	}
	err := logx.InitLogger(lc)
	if err != nil {
		fmt.Println(err)
	}

	r.Use(GinLogger(), GinRecovery(false))
	r.GET("./test", func(c *gin.Context) {
		panic("test")
		// 调用内核的Sync方法，刷新所有缓冲的日志条目。应用程序应该注意在退出之前调用Sync。
		//defer func(logger *zap.SugaredLogger) {
		//	err := logger.Sync()
		//	if err != nil {
		//
		//	}
		//}(zap.L().Sugar())
	})
	err = r.Run()
	if err != nil {
		return
	}

}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path      // 请求路径 eg: /test
		query := c.Request.URL.RawQuery //query类型的请求参数：?name=1&password=2
		// 挂起当前中间件，执行下一个中间件
		c.Next()

		cost := time.Since(start)

		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),                                 // 状态码 eg: 200
			zap.String("method", c.Request.Method),                               // 请求方法类型 eg: GET
			zap.String("path", path),                                             // 请求路径 eg: /test
			zap.String("query", query),                                           // 请求参数 eg: name=1&password=2
			zap.String("ip", c.ClientIP()),                                       // 返回真实的客户端IP eg: ::1（这个就是本机IP，ipv6地址）
			zap.String("user-agent", c.Request.UserAgent()),                      // 返回客户端的用户代理。 eg: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (HTML, like Gecko) Chrome/102.0.0.0 Safari/537.36
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()), // 返回Errors 切片中ErrorTypePrivate类型的错误
			zap.Duration("cost", cost),                                           // 返回花费时间
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 检查断开的连接，因为它不是保证紧急堆栈跟踪的真正条件。
				var brokenPipe bool
				// OpError 是 net 包中的函数通常返回的错误类型。它描述了错误的操作、网络类型和地址。
				if ne, ok := (err).(*net.OpError); ok {
					// SyscallError 记录来自特定系统调用的错误。
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") {
							brokenPipe = true
						}
					}
				}

				// DumpRequest 以 HTTP/1.x 连线形式返回给定的请求
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 如果连接死了，我们就不能给它写状态
					c.Error(err.(error))
					c.Abort() // 终止该中间件
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())), // 返回调用它的goroutine的格式化堆栈跟踪。
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 调用 `Abort()` 并使用指定的状态代码写入标头。
				// StatusInternalServerError:500
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
