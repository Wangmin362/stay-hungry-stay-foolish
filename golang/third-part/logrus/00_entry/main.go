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

	// 创建一个 Entry 对象
	entry := logger.WithFields(logrus.Fields{
		"module": "main",
	})

	// 使用 Entry 对象记录日志
	entry.Info("This is an information message")
	entry.Warn("This is a warning message")

	// 添加额外的字段到 Entry 对象
	entry = entry.WithField("user", "john")

	// 使用更新后的 Entry 对象记录日志
	entry.Error("This is an error message")

	// 在不同的位置创建 Entry 对象并记录日志
	logger.WithField("module", "submodule").Debug("Debug message from submodule")
}
