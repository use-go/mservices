package logger

import (
	"os"

	"github.com/micro/micro/v3/service/logger"
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

	DefaultLogger = logger.NewHelper(logger.NewLogger(logger.WithLevel(lvl), logger.WithCallerSkipCount(DefaultCallerSkipCount)))
}

func WithCallerSkipCount(c int) logger.Option {
	return logger.WithCallerSkipCount(c)
}
