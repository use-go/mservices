package http

import (
	"fmt"
	"math/rand"
	"net/http"

	"comm/errors"

	"github.com/2637309949/micro/v3/service/registry"
	"github.com/2637309949/micro/v3/util/selector"
)

type roundTripper struct {
	opts Options
	rt   http.RoundTripper
	st   selector.Selector
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	addr, err := r.getService(req)
	if err != nil {
		return nil, err
	}

	req.URL.Host = addr
	w, err := r.rt.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (r *roundTripper) getService(req *http.Request) (string, error) {
	var nodes []*registry.Node
	if r.opts.s != nil {
		nodes = r.opts.s.Nodes
	} else if r.opts.Router != nil {
		s, err := r.opts.Router.Route(req)
		if err != nil {
			return "", err
		}
		for _, srv := range s.Services {
			nodes = append(nodes, srv.Nodes...)
		}
	} else {
		// we have no way of routing the request
		return "", errors.InternalServerError("no route found")
	}

	// select a random node
	if len(nodes) == 0 {
		return "", errors.InternalServerError("no route found")
	}
	node := nodes[rand.Int()%len(nodes)]
	return fmt.Sprintf("http://%s", node.Address), nil
}
