package logger

import (
	"io"
	"os"

	"github.com/2637309949/micro/v3/service/logger"
)

var (
	DefaultCallerSkipCount int = 2
	DefaultLogger          logger.Logger
)

func init() {
	lvl, err := logger.GetLevel(os.Getenv("MICRO_LOG_LEVEL"))
	if err != nil {
		lvl = logger.InfoLevel
	}
	opts := []logger.Option{}
	opts = append(opts, logger.WithLevel(lvl))
	opts = append(opts, logger.WithCallerSkipCount(DefaultCallerSkipCount))
	logger.Init(opts...)
	DefaultLogger = logger.DefaultLogger
}

func WithCallerSkipCount(c int) logger.Option {
	return logger.WithCallerSkipCount(c)
}

func WithOutput(out io.Writer) logger.Option {
	return logger.WithOutput(out)
}

func WithFields(fields map[string]interface{}) logger.Option {
	return logger.WithFields(fields)
}
