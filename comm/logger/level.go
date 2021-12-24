package logger

import (
	"os"

	"github.com/2637309949/micro/v3/service/logger"
)

type Level int8

const (
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel Level = iota - 2
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// InfoLevel is the default logging priority.
	// General operational entries about what's going on inside the application.
	InfoLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	ErrorLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. highest level of severity.
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "trace"
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	}
	return ""
}

// Enabled returns true if the given level is at or above this level.
func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

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

// Returns true if the given level is at or lower the current logger level
func V(lvl Level, log logger.Logger) bool {
	l := DefaultLogger
	if log != nil {
		l = log
	}
	return l.Options().Level <= logger.Level(lvl)
}

func Init(opts ...logger.Option) {
	DefaultLogger.Init(opts...)
}
