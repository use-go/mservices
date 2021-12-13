package web

import (
	bts "bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/micro/micro/v3/service/api/handler"
	"github.com/micro/micro/v3/service/api/resolver"
	"github.com/micro/micro/v3/service/api/resolver/subdomain"
	cors "github.com/micro/micro/v3/service/api/server/http"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/debug/trace"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/util/helper"
)

type rpcRequest struct {
	Service  string
	Endpoint string
	Method   string
	Address  string
	Request  interface{}
}

type rpcHandler struct {
	resolver resolver.Resolver
	client   client.Client
}

func (h *rpcHandler) String() string {
	return "internal/rpc"
}

// see https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225
func jsonMarshal(ctx context.Context, t interface{}) ([]byte, error) {
	buffer := &bts.Buffer{}
	traceID, _, _ := trace.FromContext(ctx)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	rsp := bts.TrimRight(buffer.Bytes(), "\n")
	if strings.HasPrefix(string(rsp), "{") {
		rsp = []byte(strings.Replace(
			strings.Replace(string(rsp), "{", "{\"code\": 200,", 1),
			"{",
			"{\"request_id\": \""+traceID+"\",", 1),
		)
	}
	return rsp, err
}

// ServeHTTP passes on a JSON or form encoded RPC request to a service.
func (h *rpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		cors.SetHeaders(w, r)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	badRequest := func(description string) {
		e := errors.BadRequest("go.micro.rpc", description)
		w.WriteHeader(400)
		w.Write([]byte(e.Error()))
	}

	var service, endpoint, address string
	var request interface{}

	// response content type
	w.Header().Set("Content-Type", "application/json")

	ct := r.Header.Get("Content-Type")

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	switch ct {
	case "application/json":
		var rpcReq rpcRequest

		d := json.NewDecoder(r.Body)
		d.UseNumber()

		if err := d.Decode(&rpcReq); err != nil {
			badRequest(err.Error())
			return
		}

		service = rpcReq.Service
		endpoint = rpcReq.Endpoint
		address = rpcReq.Address
		request = rpcReq.Request
		if len(endpoint) == 0 {
			endpoint = rpcReq.Method
		}

		// JSON as string
		if req, ok := rpcReq.Request.(string); ok {
			d := json.NewDecoder(strings.NewReader(req))
			d.UseNumber()

			if err := d.Decode(&request); err != nil {
				badRequest("error decoding request string: " + err.Error())
				return
			}
		}
	default:
		r.ParseForm()
		service = r.Form.Get("service")
		endpoint = r.Form.Get("endpoint")
		address = r.Form.Get("address")
		if len(endpoint) == 0 {
			endpoint = r.Form.Get("method")
		}

		d := json.NewDecoder(strings.NewReader(r.Form.Get("request")))
		d.UseNumber()

		if err := d.Decode(&request); err != nil {
			badRequest("error decoding request string: " + err.Error())
			return
		}
	}

	if len(service) == 0 {
		badRequest("invalid service")
		return
	}

	if len(endpoint) == 0 {
		badRequest("invalid endpoint")
		return
	}

	// create request/response
	var response json.RawMessage
	var err error
	req := client.NewRequest(service, endpoint, request, client.WithContentType("application/json"))

	// create context
	ctx := helper.RequestToContext(r)
	var opts []client.CallOption

	timeout, _ := strconv.Atoi(r.Header.Get("Timeout"))
	// set timeout
	if timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(timeout)*time.Second))
	}

	// remote call
	if len(address) > 0 {
		opts = append(opts, client.WithAddress(address))
	}

	// since services can be running in many domains, we'll use the resolver to determine the domain
	// which should be used on the call
	if resolver, ok := h.resolver.(*subdomain.Resolver); ok {
		if dom := resolver.Domain(r); len(dom) > 0 {
			opts = append(opts, client.WithNetwork(dom))
		}
	}

	// remote call
	err = h.client.Call(ctx, req, &response, opts...)
	if err != nil {
		ce := errors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = "go.micro.rpc"
			ce.Status = http.StatusText(500)
			ce.Detail = "error during request: " + ce.Detail
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Code))
		}
		w.Write([]byte(ce.Error()))
		return
	}

	b, _ := jsonMarshal(ctx, response)
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}

// NewRPCHandler returns an initialized RPC handler
func NewRPCHandler(r resolver.Resolver, c client.Client) handler.Handler {
	return &rpcHandler{r, c}
}
