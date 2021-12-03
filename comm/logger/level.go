package logger

import (
	"os"

	"github.com/micro/micro/v3/service/logger"
)

func Info(args ...interface{}) {
	DefaultLogger.Log(logger.InfoLevel, args...)
}

func Infof(template string, args ...interface{}) {
	DefaultLogger.Logf(logger.InfoLevel, template, args...)
}

func Trace(args ...interface{}) {
	DefaultLogger.Log(logger.TraceLevel, args...)
}

func Tracef(template string, args ...interface{}) {
	DefaultLogger.Logf(logger.TraceLevel, template, args...)
}

func Debug(args ...interface{}) {
	DefaultLogger.Log(logger.DebugLevel, args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultLogger.Logf(logger.DebugLevel, template, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.Log(logger.WarnLevel, args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultLogger.Logf(logger.WarnLevel, template, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Log(logger.ErrorLevel, args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultLogger.Logf(logger.ErrorLevel, template, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Log(logger.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Logf(logger.FatalLevel, template, args...)
	os.Exit(1)
}

func Init(opts ...logger.Option) {
	DefaultLogger.Init(opts...)
}
