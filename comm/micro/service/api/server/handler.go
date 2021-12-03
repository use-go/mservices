package api

import (
	"net/http"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/api/handler"
	"github.com/micro/micro/v3/service/api/handler/event"
	"github.com/micro/micro/v3/service/api/router"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/errors"

	// TODO: only import handler package
	aapi "github.com/micro/micro/v3/service/api/handler/api"
	ahttp "github.com/micro/micro/v3/service/api/handler/http"
	arpc "github.com/micro/micro/v3/service/api/handler/rpc"
	aweb "github.com/micro/micro/v3/service/api/handler/web"
)

type metaHandler struct {
	c  client.Client
	r  router.Router
	ns string
}

func (m *metaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service, err := m.r.Route(r)
	if err != nil {
		er := errors.InternalServerError(m.ns, err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(er.Error()))
		return
	}

	// TODO: don't do this ffs
	switch service.Endpoint.Handler {
	// web socket handler
	case aweb.Handler:
		aweb.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// proxy handler
	case ahttp.Handler:
		ahttp.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// rpcx handler
	case arpc.Handler:
		arpc.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// event handler
	case event.Handler:
		ev := event.NewHandler(
			handler.WithNamespace(m.ns),
			handler.WithClient(m.c),
		)
		ev.ServeHTTP(w, r)
	// api handler
	case aapi.Handler:
		aapi.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	// default handler: rpc
	default:
		arpc.WithService(service, handler.WithClient(m.c)).ServeHTTP(w, r)
	}
}

// Meta is a http.Handler that routes based on endpoint metadata
func Meta(s *service.Service, r router.Router, ns string) http.Handler {
	return &metaHandler{
		c:  s.Client(),
		r:  r,
		ns: ns,
	}
}
