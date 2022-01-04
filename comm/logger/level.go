package logger

import (
	"context"
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

func AttachFields(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{"trace": ExtractTraceID(ctx)}
}

func Info(ctx context.Context, args ...interface{}) {
	Fields(AttachFields(ctx)).Log(logger.InfoLevel, args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	Fields(AttachFields(ctx)).Logf(logger.InfoLevel, template, args...)
}

func Trace(ctx context.Context, args ...interface{}) {
	Fields(AttachFields(ctx)).Log(logger.TraceLevel, args...)
}

func Tracef(ctx context.Context, template string, args ...interface{}) {
	Fields(AttachFields(ctx)).Logf(logger.TraceLevel, template, args...)
}

func Debug(ctx context.Context, args ...interface{}) {
	Fields(AttachFields(ctx)).Log(logger.DebugLevel, args...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	Fields(AttachFields(ctx)).Logf(logger.DebugLevel, template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	Fields(AttachFields(ctx)).Log(logger.WarnLevel, args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	Fields(AttachFields(ctx)).Logf(logger.WarnLevel, template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	Fields(AttachFields(ctx)).Log(logger.ErrorLevel, args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	Fields(AttachFields(ctx)).Logf(logger.ErrorLevel, template, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	Fields(AttachFields(ctx)).Log(logger.FatalLevel, args...)
	os.Exit(1)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	Fields(AttachFields(ctx)).Logf(logger.FatalLevel, template, args...)
	os.Exit(1)
}

func Fields(fields map[string]interface{}) logger.Logger {
	return DefaultLogger.Fields(fields)
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
