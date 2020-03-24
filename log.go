package log

import (
	"io"
	"os"
)

var logger *Logger

func init() {
	logger = New(os.Stdout, "", Ldatetime|Lfile)
}

func SetFlags(flag int) {
	logger.SetFlags(flag)
}

func SetLevel(level Level) {
	logger.SetLevel(level)
}

func SetPrefix(prefix string) {
	logger.SetPrefix(prefix)
}

func SetOutput(out io.Writer) {
	logger.SetOutput(out)
}
func Debugf(format string, args ...interface{}) {
	logger.logf(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.logf(InfoLevel, format, args...)
}

func Printf(format string, args ...interface{}) {
	logger.logf(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.logf(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.logf(ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.logf(FatalLevel, format, args...)
	os.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	logger.logf(PanicLevel, format, args...)
	os.Exit(1)
}

func Debug(args ...interface{}) {
	logger.log(DebugLevel, args...)
}

func Info(args ...interface{}) {
	logger.log(InfoLevel, args...)
}

func Print(args ...interface{}) {
	logger.log(InfoLevel, args...)
}

func Warn(args ...interface{}) {
	logger.log(WarnLevel, args...)
}

func Error(args ...interface{}) {
	logger.log(ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	logger.log(FatalLevel, args...)
	os.Exit(1)
}

func Panic(args ...interface{}) {
	logger.log(PanicLevel, args...)
	os.Exit(1)
}

func Println(args ...interface{}) {
	logger.log(InfoLevel, args...)
}