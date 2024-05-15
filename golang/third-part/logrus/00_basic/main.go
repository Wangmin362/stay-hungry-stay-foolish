package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	// 创建一个新的 logger 实例
	logger := logrus.New()

	// 设置日志级别为 InfoLevel
	logger.SetLevel(logrus.InfoLevel)

	// 添加一个标准输出的 Hook
	logger.SetOutput(os.Stdout)

	// 记录一条信息日志
	logger.Info("This is an information message")

	// 记录一条警告日志
	logger.Warn("This is a warning message")

	// 记录一条错误日志
	logger.Error("This is an error message")

	// 记录一条严重错误日志
	logger.Fatal("This is a fatal error message")

	// 记录一条调试日志，但是因为日志级别为 InfoLevel，所以不会被记录
	logger.Debug("This is a debug message")

	// 添加字段到日志记录中
	logger.WithFields(logrus.Fields{
		"user": "john",
		"age":  30,
	}).Info("Additional fields")

	// 通过设置环境变量来改变日志级别
	os.Setenv("LOG_LEVEL", "debug")
	// 重新读取环境变量并设置日志级别
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.Fatalf("Invalid log level: %v", err)
	}
	logger.SetLevel(logLevel)

	// 现在调试日志会被记录
	logger.Debug("This is a debug message")
}
