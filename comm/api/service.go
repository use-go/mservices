package api

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"comm/logger"

	mhttp "comm/api/http"

	meta "github.com/2637309949/micro/v3/service/context/metadata"
	"github.com/2637309949/micro/v3/service/registry"
	maddr "github.com/2637309949/micro/v3/util/addr"
	mnet "github.com/2637309949/micro/v3/util/net"
	mls "github.com/2637309949/micro/v3/util/tls"
)

type Service struct {
	opts Options

	mux *http.ServeMux
	srv *registry.Service

	sync.Mutex
	running bool
	static  bool
	exit    chan chan error
}

func newService(opts ...Option) *Service {
	options := newOptions(opts...)
	s := &Service{
		opts:   options,
		mux:    http.NewServeMux(),
		static: true,
	}
	s.srv = s.genSrv()
	return s
}

func (s *Service) genSrv() *registry.Service {
	// default host:port
	parts := strings.Split(s.opts.Address, ":")
	host := strings.Join(parts[:len(parts)-1], ":")
	port, _ := strconv.Atoi(parts[len(parts)-1])

	// check the advertise address first
	// if it exists then use it, otherwise
	// use the address
	if len(s.opts.Advertise) > 0 {
		parts = strings.Split(s.opts.Advertise, ":")

		// we have host:port
		if len(parts) > 1 {
			// set the host
			host = strings.Join(parts[:len(parts)-1], ":")

			// get the port
			if aport, _ := strconv.Atoi(parts[len(parts)-1]); aport > 0 {
				port = aport
			}
		} else {
			host = parts[0]
		}
	}

	addr, err := maddr.Extract(host)
	if err != nil {
		// best effort localhost
		addr = "127.0.0.1"
	}
	cmd := make(meta.Metadata, 1)
	cmd["protocol"] = "http"
	md := meta.Merge(s.opts.Metadata, cmd, false)

	return &registry.Service{
		Name:    s.opts.Name,
		Version: s.opts.Version,
		Nodes: []*registry.Node{{
			Id:       s.opts.Id,
			Address:  fmt.Sprintf("%s:%d", addr, port),
			Metadata: md,
		}},
	}
}

func (s *Service) run(exit chan bool) {
	if s.opts.RegisterInterval <= time.Duration(0) {
		return
	}

	t := time.NewTicker(s.opts.RegisterInterval)

	for {
		select {
		case <-t.C:
			s.register()
		case <-exit:
			t.Stop()
			return
		}
	}
}

func (s *Service) register() error {
	if s.srv == nil {
		return nil
	}
	// default to service registry
	r := registry.DefaultRegistry
	// switch to option if specified
	if s.opts.Registry != nil {
		r = s.opts.Registry
	}

	// service node need modify, node address maybe changed
	srv := s.genSrv()
	srv.Endpoints = s.srv.Endpoints
	s.srv = srv
	return r.Register(s.srv, registry.RegisterTTL(s.opts.RegisterTTL))
}

func (s *Service) deregister() error {
	if s.srv == nil {
		return nil
	}
	// default to service registry
	r := registry.DefaultRegistry
	// switch to option if specified
	if s.opts.Registry != nil {
		r = s.opts.Registry
	}
	return r.Deregister(s.srv)
}

func (s *Service) start() error {
	s.Lock()
	defer s.Unlock()

	if s.running {
		return nil
	}

	l, err := s.listen("tcp", s.opts.Address)
	if err != nil {
		return err
	}

	s.opts.Address = l.Addr().String()
	srv := s.genSrv()
	srv.Endpoints = s.srv.Endpoints
	s.srv = srv

	var h http.Handler

	if s.opts.Handler != nil {
		h = s.opts.Handler
	} else {
		h = s.mux
		var r sync.Once

		// register the html dir
		r.Do(func() {
			// static dir
			static := s.opts.StaticDir
			if s.opts.StaticDir[0] != '/' {
				dir, _ := os.Getwd()
				static = filepath.Join(dir, static)
			}

			// set static if no / handler is registered
			if s.static {
				_, err := os.Stat(static)
				if err == nil {
					logger.Infof("Enabling static file serving from %s", static)
					s.mux.Handle("/", http.FileServer(http.Dir(static)))
				}
			}
		})
	}

	for _, fn := range s.opts.BeforeStart {
		if err := fn(); err != nil {
			return err
		}
	}

	var httpSrv *http.Server
	if s.opts.Server != nil {
		httpSrv = s.opts.Server
	} else {
		httpSrv = &http.Server{}
	}

	httpSrv.Handler = h

	go httpSrv.Serve(l)

	for _, fn := range s.opts.AfterStart {
		if err := fn(); err != nil {
			return err
		}
	}

	s.exit = make(chan chan error, 1)
	s.running = true

	go func() {
		ch := <-s.exit
		ch <- l.Close()
	}()

	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("HTTP API Listening on %s", l.Addr().String())
	}

	logger.Infof("HTTP API Listening on %v", l.Addr().String())
	return nil
}

func (s *Service) stop() error {
	s.Lock()
	defer s.Unlock()

	if !s.running {
		return nil
	}

	for _, fn := range s.opts.BeforeStop {
		if err := fn(); err != nil {
			return err
		}
	}

	ch := make(chan error, 1)
	s.exit <- ch
	s.running = false

	logger.Info("Stopping")

	for _, fn := range s.opts.AfterStop {
		if err := fn(); err != nil {
			if chErr := <-ch; chErr != nil {
				return chErr
			}
			return err
		}
	}

	return <-ch
}

func (s *Service) Client() *http.Client {
	rt := mhttp.NewRoundTripper(
		mhttp.WithRegistry(registry.DefaultRegistry),
		mhttp.WithService(s.srv),
	)
	return &http.Client{
		Transport: rt,
	}
}

func (s *Service) Handle(pattern string, handler http.Handler) {
	var seen bool
	for _, ep := range s.srv.Endpoints {
		if ep.Name == pattern {
			seen = true
			break
		}
	}

	// if its unseen then add an endpoint
	if !seen {
		s.srv.Endpoints = append(s.srv.Endpoints, &registry.Endpoint{
			Name: pattern,
		})
	}

	// disable static serving
	if pattern == "/" {
		s.Lock()
		s.static = false
		s.Unlock()
	}

	// register the handler
	pattern = "/" + s.opts.Name + pattern
	s.mux.Handle(pattern, handler)
}

func (s *Service) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	var seen bool
	for _, ep := range s.srv.Endpoints {
		if ep.Name == pattern {
			seen = true
			break
		}
	}
	if !seen {
		s.srv.Endpoints = append(s.srv.Endpoints, &registry.Endpoint{
			Name: pattern,
		})
	}

	pattern = "/" + s.opts.Name + pattern
	s.mux.HandleFunc(pattern, handler)
}

func (s *Service) Run() error {
	if err := s.start(); err != nil {
		return err
	}

	if err := s.register(); err != nil {
		return err
	}

	// start reg loop
	ex := make(chan bool)
	go s.run(ex)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	select {
	// wait on kill signal
	case sig := <-ch:
		logger.Infof("Received signal %s", sig)
	// wait on context cancel
	case <-s.opts.Context.Done():
		logger.Infof("Received context shutdown")
	}

	// exit reg loop
	close(ex)

	if err := s.deregister(); err != nil {
		return err
	}

	return s.stop()
}

// Options returns the options for the given service
func (s *Service) Options() Options {
	return s.opts
}

func (s *Service) listen(network, addr string) (net.Listener, error) {
	var l net.Listener
	var err error

	// TODO: support use of listen options
	if s.opts.Secure || s.opts.TLSConfig != nil {
		config := s.opts.TLSConfig

		fn := func(addr string) (net.Listener, error) {
			if config == nil {
				hosts := []string{addr}

				// check if its a valid host:port
				if host, _, err := net.SplitHostPort(addr); err == nil {
					if len(host) == 0 {
						hosts = maddr.IPs()
					} else {
						hosts = []string{host}
					}
				}

				// generate a certificate
				cert, err := mls.Certificate(hosts...)
				if err != nil {
					return nil, err
				}
				config = &tls.Config{Certificates: []tls.Certificate{cert}}
			}
			return tls.Listen(network, addr, config)
		}

		l, err = mnet.Listen(addr, fn)
	} else {
		fn := func(addr string) (net.Listener, error) {
			return net.Listen(network, addr)
		}

		l, err = mnet.Listen(addr, fn)
	}

	if err != nil {
		return nil, err
	}

	return l, nil
}
