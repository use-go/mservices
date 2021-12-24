package web

import (
	"comm/logger"
	"net/http"
	"time"

	"github.com/2637309949/micro/v3/service"
	"github.com/2637309949/micro/v3/service/api"
	apiAuth "github.com/2637309949/micro/v3/service/api/auth"
	"github.com/2637309949/micro/v3/service/api/resolver"
	httpapi "github.com/2637309949/micro/v3/service/api/server/http"
	"github.com/2637309949/micro/v3/service/registry"
	"github.com/2637309949/micro/v3/service/router"
	regRouter "github.com/2637309949/micro/v3/service/router/registry"
	"github.com/2637309949/micro/v3/service/web"
	cx "github.com/2637309949/micro/v3/util/ctx"
	"github.com/gorilla/mux"
)

type Service struct {
	*service.Service
	mux *Mux
}

func (s *Service) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return s.mux.HandleFunc(path, f)
}

func (s *Service) Run(address string) error {
	var apis []api.Option
	var h http.Handler = s.mux
	server := httpapi.NewServer(address)
	aw := apiAuth.Wrapper(s.mux.resolver, Namespace)
	server.Init(apis...)
	server.Handle("/", aw(h))

	if err := server.Start(); err != nil {
		logger.Error(err)
		return err
	}

	// Run service
	if err := s.Service.Run(); err != nil {
		logger.Error(err)
		return err
	}

	if err := server.Stop(); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func New(opts ...service.Option) *Service {
	opts = append(opts, service.Version("latest"))
	opts = append(opts, service.RegisterTTL(5*time.Second))
	opts = append(opts, service.RegisterInterval(5*time.Second))
	opts = append(opts, service.Metadata(map[string]string{"protocol": "http"}))
	opts = append(opts, debugWrapper)
	mux := Mux{
		Router: mux.NewRouter(),
		registry: &reg{
			Registry: registry.DefaultRegistry,
		},
		resolver: &web.WebResolver{
			Router:  regRouter.NewRouter(router.Registry(registry.DefaultRegistry)),
			Options: resolver.NewOptions(resolver.WithServicePrefix(Namespace)),
		},
	}
	mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(cx.FromRequest(r))
			next.ServeHTTP(w, r)
		})
	})
	return &Service{
		mux:     &mux,
		Service: service.New(opts...),
	}
}

func Name(n string) service.Option {
	return service.Name(n)
}
