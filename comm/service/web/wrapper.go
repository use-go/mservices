package web

import (
	"bytes"
	"comm/errors"
	"comm/trace"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	sAuth "github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/util/auth"
	cx "github.com/2637309949/micro/v3/util/ctx"
	uhttp "github.com/2637309949/micro/v3/util/http"
)

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	// reduce mem
	if r.ResponseWriter.Header().Get("Content-Type") == "application/json" {
		r.body.Write(b)
	}
	return r.ResponseWriter.Write(b)
}

var (
	debugWrapper = func(call func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(res http.ResponseWriter, req *http.Request) {
			req = req.WithContext(cx.FromRequest(req))
			bodyBytes := []byte("{}")
			if req.Header.Get("Content-Type") == "application/json" {
				bodyBytes, _ := ioutil.ReadAll(req.Body)
				req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
			w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: res}
			defer func() {
				trace.Debug(cx.FromRequest(req), req.URL.Path, bodyBytes, w.body.Bytes())()
			}()
			call(w, req)
		}
	}
	authWrapper = func(call func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
		return func(res http.ResponseWriter, req *http.Request) {
			var token string
			if header := req.Header.Get("Authorization"); len(header) > 0 {
				// Ensure the correct scheme is being used
				switch {
				case strings.HasPrefix(header, auth.BearerScheme):
					// Strip the bearer scheme prefix
					token = strings.TrimPrefix(header, auth.BearerScheme)
				case strings.HasPrefix(header, "Basic "):
					b, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(header, "Basic "))
					if err != nil {
						uhttp.WriteError(res, req, errors.Unauthorized("invalid authorization header. Incorrect format"))
						return
					}
					parts := strings.SplitN(string(b), ":", 2)
					if len(parts) != 2 {
						uhttp.WriteError(res, req, errors.Unauthorized("invalid authorization header. Incorrect format"))
						return
					}
					token = parts[1]
				default:
				}
			}
			// Determine the namespace
			ns := sAuth.DefaultAuth.Options().Issuer

			var acc *sAuth.Account
			if a, err := sAuth.Inspect(token); err == nil {
				ctx := sAuth.ContextWithAccount(req.Context(), a)
				*req = *req.Clone(ctx)
				acc = a
			}

			// construct the resource
			re := &sAuth.Resource{
				Type:     "service",
				Name:     "",
				Endpoint: "",
			}
			// Verify the caller has access to the resource.
			err := sAuth.Verify(acc, re, sAuth.VerifyNamespace(ns))
			if err == sAuth.ErrForbidden && acc != nil {
				uhttp.WriteError(res, req, errors.Forbidden("Forbidden call made to %v:%v by %v", re.Name, re.Endpoint, acc.ID))
				return
			} else if err == sAuth.ErrForbidden {
				uhttp.WriteError(res, req, errors.Unauthorized("Unauthorized call made to %v:%v", re.Name, re.Endpoint))
				return
			} else if err != nil {
				uhttp.WriteError(res, req, errors.InternalServerError("Error authorizing request: %v", err))
				return
			}
			call(res, req)
		}
	}
)
