package main

import (
	"errors"
	"go-log/logx"
	"go.uber.org/zap"
)

func main() {

	conf := logx.LogConfig{
		LogLevel:          "debug",    // 输出日志级别 "debug" "info" "warn" "error"
		LogFormat:         "json",     // 输出日志格式  json
		LogPath:           "./log",    // 输出日志文件位置
		LogFileName:       "test.log", // 输出日志文件名称
		LogFileMaxSize:    1,          // 输出单个日志文件大小，单位MB
		LogFileMaxBackups: 10,         // 输出最大日志备份个数
		LogMaxAge:         1000,       // 日志保留时间，单位: 天 (day)
		LogCompress:       false,      // 是否压缩日志
		LogStdout:         false,      // 是否输出到控制台
	}
	// 2. 初始化log
	if err := logx.InitLogger(conf); err != nil {
		panic(err)
	}

	zap.S().Debugf("测试 Debugf 用法：%s", "111") // logger Debugf 用法
	zap.S().Errorf("测试 Errorf 用法：%s", "111") // logger Errorf 用法
	zap.S().Warnf("测试 Warnf 用法：%s", "111")   // logger Warnf 用法
	zap.S().Infof("测试 Infof 用法：%s, %d, %v, %f", "111", 1111, errors.New("collector returned no data"), 3333.33)
	// logger With 用法
	logger := zap.S().With("collector", "cpu", "name", "主机")
	logger.Infof("测试 (With + Infof) 用法：%s", "测试")
	zap.S().Errorf("测试 Errorf 用法：%s", "111")

}
