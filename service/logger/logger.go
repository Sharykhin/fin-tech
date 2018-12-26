package logger

import (
	"log"

	"github.com/Sharykhin/fin-tech/contract"
	"github.com/Sharykhin/fin-tech/service/logger/std"
)

var Logger contract.Logger = std.NewStdLogger()

const (
	ERROR     = "ERROR"
	INFO      = "INFO"
	EMERGENCY = "EMERGENCY"
)

func Error(format string, v ...interface{}) {
	Logger.Error(format, v)
}

func Info(format string, v ...interface{}) {
	Logger.Info(format, v)
}

func Emergency(format string, v ...interface{}) {
	Logger.Emergency(format, v)
}

func LOG(method string, format string, v ...interface{}) {
	switch method {
	case INFO:
		Logger.Info(format, v)
	case ERROR:
		Logger.Error(format, v)
	case EMERGENCY:
		Logger.Emergency(format, v)
	default:
		log.Panicf("method LOG was called with invalid parameters")
	}
}
