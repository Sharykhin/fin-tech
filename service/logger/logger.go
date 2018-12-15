package logger

import (
	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/service/logger/std"
)

var Logger contract.Logger = std.NewStdLogger()

func Error(format string, v ...interface{}) {
	Logger.Error(format, v)
}

func Info(format string, v ...interface{}) {
	Logger.Info(format, v)
}

func Emergency(format string, v ...interface{}) {
	Logger.Emergency(format, v)
}
