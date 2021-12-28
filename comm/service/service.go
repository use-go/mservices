package service

import (
	"comm/logger"
	"io"
	"log"
	"os"
	"time"

	"github.com/2637309949/micro/v3/service"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type Service struct {
	*service.Service
}

func (s *Service) Run() error {
	err := s.streamOutput()
	if err != nil {
		return nil
	}
	err = s.Service.Run()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) streamOutput() error {
	// make the logs directory
	fp := logFile(s.Options().Name)
	out := os.Stdout
	rotate, err := rotatelogs.New(
		fp+"/%Y%m%d%H",
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		return err
	}

	mw := io.MultiWriter(out, rotate)

	log.SetOutput(mw)
	logger.Init(logger.WithOutput(mw))

	// all writes to stdout,stderr will go through pipe instead (fmt.print, log)
	r, w, err := os.Pipe()
	os.Stdout = w
	os.Stderr = w
	go io.Copy(mw, r)

	return err
}

func newOptions(opts ...service.Option) []service.Option {
	opts = append(opts, service.Version("latest"))
	opts = append(opts, service.RegisterTTL(5*time.Second))
	opts = append(opts, service.RegisterInterval(5*time.Second))
	opts = append(opts, debugWrapper)
	return opts
}

func New(opts ...service.Option) *Service {
	opts = newOptions(opts...)
	srv := service.New(opts...)
	return &Service{
		srv,
	}
}

func Name(n string) service.Option {
	return service.Name(n)
}
