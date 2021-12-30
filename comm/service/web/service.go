package web

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
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

	mhttp "comm/service/web/http"

	apiAuth "github.com/2637309949/micro/v3/service/api/auth"
	"github.com/2637309949/micro/v3/service/api/resolver"
	meta "github.com/2637309949/micro/v3/service/context/metadata"
	"github.com/2637309949/micro/v3/service/registry"
	"github.com/2637309949/micro/v3/service/router"
	regRouter "github.com/2637309949/micro/v3/service/router/registry"
	"github.com/2637309949/micro/v3/service/web"
	maddr "github.com/2637309949/micro/v3/util/addr"
	"github.com/2637309949/micro/v3/util/backoff"
	mnet "github.com/2637309949/micro/v3/util/net"
	mls "github.com/2637309949/micro/v3/util/tls"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

type Service struct {
	opts Options

	mux *http.ServeMux
	srv *registry.Service

	registered bool

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

	if s.opts.RegisterCheck == nil {
		s.opts.RegisterCheck = DefaultRegisterCheck
	}

	if s.opts.Registry == nil {
		s.opts.Registry = registry.DefaultRegistry
	}

	s.srv = s.genSrv()
	return s
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
		addr = "127.0.0.1"
	}
	cmd := make(meta.Metadata, 1)
	cmd["protocol"] = "http"
	cmd["registry"] = "service"
	cmd["server"] = "http"
	cmd["transport"] = "grpc"
	md := meta.Merge(s.opts.Metadata, cmd, false)

	rsMeta := make(meta.Metadata, 1)
	rsMeta["domain"] = Namespace
	rsMeta["handler"] = "http"
	return &registry.Service{
		Name:     s.opts.Name,
		Version:  s.opts.Version,
		Metadata: rsMeta,
		Nodes: []*registry.Node{{
			Id:       s.opts.Name + "-" + s.opts.Id,
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
			registered := s.registered
			rerr := s.opts.RegisterCheck(s.opts.Context)
			if rerr != nil && registered {
				if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
					logger.Errorf("Server %s-%s register check error: %s, deregister it", s.opts.Name, s.srv.Nodes[0].Id, rerr)
				}
				// deregister self in case of error
				if err := s.deregister(); err != nil {
					if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
						logger.Errorf("Server %s-%s deregister error: %s", s.opts.Name, s.srv.Nodes[0].Id, err)
					}
				}
			} else if rerr != nil && !registered {
				if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
					logger.Errorf("Server %s-%s register check error: %s", s.opts.Name, s.srv.Nodes[0].Id, rerr)
				}
				continue
			}
			if err := s.register(); err != nil {
				if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
					logger.Errorf("Server %s-%s register error: %s", s.opts.Name, s.srv.Nodes[0].Id, err)
				}
			}
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

	regFunc := func(service *registry.Service) error {
		// create registry options
		rOpts := []registry.RegisterOption{
			registry.RegisterTTL(s.opts.RegisterTTL),
		}

		var regErr error

		for i := 0; i < 3; i++ {
			// attempt to register
			if err := s.opts.Registry.Register(service, rOpts...); err != nil {
				// set the error
				regErr = err
				// backoff then retry
				time.Sleep(backoff.Do(i + 1))
				continue
			}
			// success so nil error
			regErr = nil
			break
		}

		return regErr
	}

	// service node need modify, node address maybe changed
	srv := s.genSrv()
	srv.Endpoints = s.srv.Endpoints
	s.srv = srv

	// get registered value
	registered := s.registered

	if !registered {
		if logger.V(logger.InfoLevel, logger.DefaultLogger) {
			logger.Infof("Registry [%s] Registering node: %s", s.opts.Registry.String(), s.srv.Nodes[0].Id)
		}
	}

	// register the service
	if err := regFunc(s.srv); err != nil {
		return err
	}

	// already registered? don't need to register subscribers
	if registered {
		return nil
	}
	s.registered = true

	return nil
}

func (s *Service) deregister() error {
	if s.srv == nil {
		return nil
	}

	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Deregistering node: %s", s.srv.Nodes[0].Id)
	}
	return s.opts.Registry.Deregister(s.srv)
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
		mhttp.WithRegistry(s.opts.Registry),
		mhttp.WithService(s.srv),
	)
	return &http.Client{
		Transport: rt,
	}
}

func (s *Service) Handle(pattern string, handler http.Handler) {

	// disable static serving
	if pattern == "/" {
		s.Lock()
		s.static = false
		s.Unlock()
	}

	// register the handler
	s.mux.Handle("/"+s.opts.Name+pattern, handler)
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

	// register the handler
	s.mux.HandleFunc("/"+s.opts.Name+pattern, s.opts.wrapper(handler))
}

func (s *Service) Run() error {
	err := s.streamOutput()
	if err != nil {
		return nil
	}

	resolver := &web.WebResolver{
		Router:  regRouter.NewRouter(router.Registry(s.opts.Registry)),
		Options: resolver.NewOptions(resolver.WithServicePrefix(Namespace)),
	}
	aw := apiAuth.Wrapper(resolver, Namespace)
	s.Handle("/", aw(s.mux))

	if logger.V(logger.InfoLevel, logger.DefaultLogger) {
		logger.Infof("Starting [service] %s", s.opts.Name)
	}

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
	case <-ch:
	// wait on context cancel
	case <-s.opts.Context.Done():
	}

	// exit reg loop
	close(ex)

	registered := s.registered
	if registered {
		if err := s.deregister(); err != nil {
			return err
		}
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
