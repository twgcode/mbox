/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2023/4/26 14:10
@Description:
*/

package sync

import (
	"fmt"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func TestMain(m *testing.M) {
	var (
		err error
	)
	config := zap.Config{
		Encoding:         "console", // 使用 ConsoleEncoder
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stdout"}, // 将日志输出到 stdout
		ErrorOutputPaths: []string{"stderr"}, // 将错误日志输出到 stderr
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			EncodeTime:    zapcore.ISO8601TimeEncoder,       // 时间格式化
			EncodeLevel:   zapcore.CapitalColorLevelEncoder, // Level 字段格式化
			EncodeCaller:  zapcore.ShortCallerEncoder,       // Caller 字段格式化
		},
	}

	logger, err = config.Build(zap.AddStacktrace(zapcore.ErrorLevel))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	logger.Info("Zap logger initialized.")
	retCode := m.Run() // 执行测试
	os.Exit(retCode)   // 退出测试
}

func TestGo(t *testing.T) {
	t.Run("Test normal case", func(t *testing.T) {
		// 定义一个简单的函数用于测试
		testFunc := func() {
			fmt.Println("Hello, world!")
		}
		// 在后台运行该函数，并等待一段时间以确保其已完成执行
		Go(testFunc, "test1", nil)
		time.Sleep(100 * time.Millisecond)
	})

	t.Run("Test panic case", func(t *testing.T) {
		// 定义一个会触发panic的函数用于测试
		testFunc := func() {
			panic("Oops, something went wrong!")
		}
		// 在后台运行该函数，并等待一段时间以确保其已完成执行
		Go(testFunc, "test2", logger)
		time.Sleep(100 * time.Millisecond)
	})
}
