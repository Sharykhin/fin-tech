package std

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func (l Logger) Error(format string, v ...interface{}) {
	l.Printf("ERROR: "+format, v...)
}

func (l Logger) Info(format string, v ...interface{}) {
	l.Printf("INFO: "+format, v...)
}

func (l Logger) Emergency(format string, v ...interface{}) {
	l.Panicf("EMERGENCY: "+format, v...)
}

// NewStdLogger returns a new instance of logger that put logs into std output
func NewStdLogger() *Logger {
	l := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	return &Logger{l}
}
