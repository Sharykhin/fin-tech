package contract

type (
	Logger interface {
		Error(format string, v ...interface{})
		Info(format string, v ...interface{})
		Emergency(format string, v ...interface{})
	}
)
