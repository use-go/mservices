package trace

import (
	"io"

	"github.com/2637309949/micro/v3/service/config"
	"github.com/2637309949/micro/v3/service/logger"
	"github.com/2637309949/micro/v3/util/opentelemetry"
	"github.com/2637309949/micro/v3/util/opentelemetry/jaeger"
)

func SetupOpentracing(serviceName string) io.Closer {
	c, _ := config.Get("jaegeraddress")
	openTracer, closer, err := jaeger.New(
		opentelemetry.WithServiceName(serviceName),
		opentelemetry.WithTraceReporterAddress(c.String("localhost:6831")),
	)
	if err != nil {
		logger.Fatalf("Error configuring opentracing: %v", err)
	}
	logger.Infof("Configured jaeger to %s", c.String("localhost:6831"))
	opentelemetry.DefaultOpenTracer = openTracer
	return closer
}
