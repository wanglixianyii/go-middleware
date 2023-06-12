package main

import (
	"fmt"
	zapLogger "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go-log/logx"
	"go.uber.org/zap"
	"time"
)

func main() {
	serviceInfo := &ServiceInfo{
		Name:    "test",
		Version: "1.0",
		Id:      "123123213",
	}
	// init logger
	logger := NewLoggerProvider(serviceInfo)

	// 日志中间件
	kratos.New(
		kratos.Logger(logger),
	)

}

type ServiceInfo struct {
	Name     string
	Version  string
	Id       string
	Metadata map[string]string
}

func NewLoggerProvider(serviceInfo *ServiceInfo) log.Logger {
	cfg := logx.LogConfig{
		LogLevel:          "debug",
		LogFileName:       fmt.Sprintf("%v.log", time.Now().Unix()),
		LogFileMaxSize:    1,
		LogFileMaxBackups: 5,
		LogMaxAge:         30,
	}
	err := logx.InitLogger(cfg)
	if err != nil {
		fmt.Println(err)
	}

	l := zapLogger.NewLogger(zap.L())
	return log.With(
		l,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}
