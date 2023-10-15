package log

type Logger interface {
	Debug(message string)
	Debugf(format string, args ...interface{})
	Info(message string)
	Infof(format string, args ...interface{})
	Error(message string)
	Errorf(format string, args ...interface{})
	Fatal(message string)
	Fatalf(format string, args ...interface{})
}

func Debug(message string) {
	Log.Debug(message)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Info(message string) {
	Log.Info(message)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Error(message string) {
	Log.Error(message)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Fatal(message string) {
	Log.Fatal(message)
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}
