// Package web is a web dashboard
package web

import (
	"comm/logger"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/2637309949/micro/v3/service/api/resolver"
	"github.com/2637309949/micro/v3/service/api/resolver/subdomain"
	httpapi "github.com/2637309949/micro/v3/service/api/server/http"
	"github.com/2637309949/micro/v3/service/auth"
	"github.com/2637309949/micro/v3/service/registry"
	"github.com/2637309949/micro/v3/util/acme"
	"github.com/fatih/camelcase"
	"github.com/gorilla/mux"
	"github.com/serenize/snaker"
)

//Meta Fields of micro web
var (
	Namespace             = "micro"
	Resolver              = "path"
	LoginURL              = "/login"
	ACMEProvider          = "autocert"
	ACMEChallengeProvider = "cloudflare"
	ACMECA                = acme.LetsEncryptProductionCA

	// Host name the web dashboard is served on
	Host, _ = os.Hostname()
	// Token cookie name
	TokenCookieName = "micro-token"
)

type Mux struct {
	*mux.Router
	// registry we use
	registry registry.Registry
	// the resolver
	resolver resolver.Resolver
}

type reg struct {
	registry.Registry
	sync.RWMutex
	lastPull time.Time
	services []*registry.Service
}

// ServeHTTP serves the web dashboard and proxies where appropriate
func (s *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check if authenticated
	if r.URL.Path != LoginURL {
		c, err := r.Cookie(TokenCookieName)
		if err != nil || c == nil {
			http.Redirect(w, r, LoginURL, 302)
			return
		}

		// check the token is valid
		token := strings.TrimPrefix(c.Value, TokenCookieName+"=")
		if len(token) == 0 {
			http.Redirect(w, r, LoginURL, 302)
			return
		}
	}

	// set defaults on the request
	if len(r.URL.Host) == 0 {
		r.URL.Host = r.Host
	}
	if len(r.URL.Scheme) == 0 {
		r.URL.Scheme = "http"
	}

	// no endpoint was set in the context, so we'll look it up. If the router returns an error we will
	// send the request to the mux which will render the web dashboard.
	s.Router.ServeHTTP(w, r)
	//_, err := s.resolver.Resolve(r)
	//if err != nil {
	//	return
	//}
}

func split(v string) string {
	parts := camelcase.Split(strings.Replace(v, ".", "", 1))
	return strings.Join(parts, " ")
}

func last(v string) string {
	parts := strings.Split(v, ".")
	return split(parts[len(parts)-1])
}

func format(v *registry.Value) string {
	if v == nil || len(v.Values) == 0 {
		return "{}"
	}
	var f []string
	for _, k := range v.Values {
		f = append(f, formatEndpoint(k, 0))
	}
	return fmt.Sprintf("{\n%s}", strings.Join(f, ""))
}

func formatEndpoint(v *registry.Value, r int) string {
	// default format is tabbed plus the value plus new line
	fparts := []string{"", "%s %s", "\n"}
	for i := 0; i < r+1; i++ {
		fparts[0] += "\t"
	}
	// its just a primitive of sorts so return
	if len(v.Values) == 0 {
		return fmt.Sprintf(strings.Join(fparts, ""), snaker.CamelToSnake(v.Name), v.Type)
	}

	// this thing has more things, it's complex
	fparts[1] += " {"

	vals := []interface{}{snaker.CamelToSnake(v.Name), v.Type}

	for _, val := range v.Values {
		fparts = append(fparts, "%s")
		vals = append(vals, formatEndpoint(val, r+1))
	}

	// at the end
	l := len(fparts) - 1
	for i := 0; i < r+1; i++ {
		fparts[l] += "\t"
	}
	fparts = append(fparts, "}\n")

	return fmt.Sprintf(strings.Join(fparts, ""), vals...)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Mux) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.render(w, r, notFoundTemplate, nil)
}

func (s *Mux) indexHandler(w http.ResponseWriter, r *http.Request) {
	httpapi.SetHeaders(w, r)

	if r.Method == "OPTIONS" {
		return
	}

	// if we're using the subdomain resolver, we want to use a custom domain
	domain := registry.DefaultDomain
	if res, ok := s.resolver.(*subdomain.Resolver); ok {
		domain = res.Domain(r)
	}

	services, err := s.registry.ListServices(registry.ListDomain(domain))
	if err != nil {
		logger.Errorf("Error listing services: %v", err)
	}

	type webService struct {
		Name string
		Link string
		Icon string // TODO: lookup icon
	}

	var webServices []webService
	for _, srv := range services {
		name := srv.Name

		if len(srv.Endpoints) == 0 {
			continue
		}

		// in the case of 3 letter things e.g m3o convert to M3O
		if len(name) <= 3 && strings.ContainsAny(name, "012345789") {
			name = strings.ToUpper(name)
		}

		webServices = append(webServices, webService{Name: name, Link: fmt.Sprintf("/%v", name)})
	}

	sort.Slice(webServices, func(i, j int) bool { return webServices[i].Name < webServices[j].Name })

	type templateData struct {
		HasWebServices bool
		WebServices    []webService
	}

	data := templateData{len(webServices) > 0, webServices}
	s.render(w, r, indexTemplate, data)
}

func (s *Mux) loginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		s.generateTokenHandler(w, req)
		return
	}

	t, err := template.New("template").Parse(loginTemplate)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "basic", map[string]interface{}{
		"foo": "bar",
	}); err != nil {
		http.Error(w, "Error occurred:"+err.Error(), http.StatusInternalServerError)
	}
}

func (s *Mux) logoutHandler(w http.ResponseWriter, req *http.Request) {
	var domain string
	if arr := strings.Split(req.Host, ":"); len(arr) > 0 {
		domain = arr[0]
	}

	http.SetCookie(w, &http.Cookie{
		Name:    TokenCookieName,
		Value:   "",
		Expires: time.Unix(0, 0),
		Domain:  domain,
		Secure:  true,
	})

	http.Redirect(w, req, "/", http.StatusFound)
}

func (s *Mux) generateTokenHandler(w http.ResponseWriter, req *http.Request) {
	renderError := func(errMsg string) {
		t, err := template.New("template").Parse(loginTemplate)
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := t.ExecuteTemplate(w, "basic", map[string]interface{}{
			"error": errMsg,
		}); err != nil {
			http.Error(w, "Error occurred:"+err.Error(), http.StatusInternalServerError)
		}
	}

	user := req.PostFormValue("username")
	if len(user) == 0 {
		renderError("Missing Username")
		return
	}

	pass := req.PostFormValue("password")
	if len(pass) == 0 {
		renderError("Missing Password")
		return
	}

	acc, err := auth.Token(
		auth.WithCredentials(user, pass),
		auth.WithTokenIssuer(Namespace),
		auth.WithExpiry(time.Hour*24*7),
	)

	if err != nil {
		renderError("Authentication failed: " + err.Error())
		return
	}

	var domain string
	if arr := strings.Split(req.Host, ":"); len(arr) > 0 {
		domain = arr[0]
	}
	http.SetCookie(w, &http.Cookie{
		Name:    TokenCookieName,
		Value:   acc.AccessToken,
		Expires: acc.Expiry,
		Domain:  domain,
		Secure:  false,
	})

	http.Redirect(w, req, "/", http.StatusFound)
}

func (s *Mux) registryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	svc := vars["name"]

	// if we're using the subdomain resolver, we want to use a custom domain
	domain := registry.DefaultDomain
	if res, ok := s.resolver.(*subdomain.Resolver); ok {
		domain = res.Domain(r)
	}

	if len(svc) > 0 {
		sv, err := s.registry.GetService(svc, registry.GetDomain(domain))
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}

		if len(sv) == 0 {
			http.Error(w, "Not found", 404)
			return
		}

		if r.Header.Get("Content-Type") == "application/json" {
			b, err := json.Marshal(map[string]interface{}{
				"services": s,
			})
			if err != nil {
				http.Error(w, "Error occurred:"+err.Error(), 500)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(b)
			return
		}

		s.render(w, r, serviceTemplate, sv)
		return
	}

	services, err := s.registry.ListServices(registry.ListDomain(domain))
	if err != nil {
		logger.Errorf("Error listing services: %v", err)
	}

	sort.Sort(sortedServices{services})

	if r.Header.Get("Content-Type") == "application/json" {
		b, err := json.Marshal(map[string]interface{}{
			"services": services,
		})
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	s.render(w, r, registryTemplate, services)
}

func (s *Mux) callHandler(w http.ResponseWriter, r *http.Request) {
	// if we're using the subdomain resolver, we want to use a custom domain
	domain := registry.DefaultDomain
	if res, ok := s.resolver.(*subdomain.Resolver); ok {
		domain = res.Domain(r)
	}

	services, err := s.registry.ListServices(registry.ListDomain(domain))
	if err != nil {
		logger.Errorf("Error listing services: %v", err)
	}

	sort.Sort(sortedServices{services})

	serviceMap := make(map[string][]*registry.Endpoint)
	for _, service := range services {
		if len(service.Endpoints) > 0 {
			serviceMap[service.Name] = service.Endpoints
			continue
		}
		// lookup the endpoints otherwise
		s, err := s.registry.GetService(service.Name, registry.GetDomain(domain))
		if err != nil {
			continue
		}
		if len(s) == 0 {
			continue
		}
		serviceMap[service.Name] = s[0].Endpoints
	}

	if r.Header.Get("Content-Type") == "application/json" {
		b, err := json.Marshal(map[string]interface{}{
			"services": services,
		})
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	s.render(w, r, callTemplate, serviceMap)
}

func (s *Mux) serviceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["service"]
	if len(name) == 0 {
		return
	}

	// if we're using the subdomain resolver, we want to use a custom domain
	domain := registry.DefaultDomain
	if res, ok := s.resolver.(*subdomain.Resolver); ok {
		domain = res.Domain(r)
	}

	services, err := s.registry.GetService(name, registry.GetDomain(domain))
	if err != nil {
		logger.Errorf("Error getting service %s: %v", name, err)
	}

	sort.Sort(sortedServices{services})

	serviceMap := make(map[string][]*registry.Endpoint)

	for _, service := range services {
		for _, endpoint := range service.Endpoints {
			service := strings.Split(endpoint.Name, ".")[0]
			if _, ok := serviceMap[service]; !ok {
				serviceMap[service] = []*registry.Endpoint{}
			}
			serviceMap[service] = append(serviceMap[service], endpoint)
		}
	}

	if r.Header.Get("Content-Type") == "application/json" {
		b, err := json.Marshal(map[string]interface{}{
			"services": services,
		})
		if err != nil {
			http.Error(w, "Error occurred:"+err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}

	s.render(w, r, webTemplate, serviceMap, templateValue{
		Key:   "Name",
		Value: name,
	})
}

type templateValue struct {
	Key   string
	Value interface{}
}

func (s *Mux) render(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}, vals ...templateValue) {
	t, err := template.New("template").Funcs(template.FuncMap{
		"Split":  split,
		"Last":   last,
		"format": format,
		"Title":  strings.Title,
		"First": func(s string) string {
			if len(s) == 0 {
				return s
			}
			return strings.Title(string(s[0]))
		},
	}).Parse(layoutTemplate)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}
	t, err = t.Parse(tmpl)
	if err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
		return
	}

	// If the user is logged in, render Account instead of Login
	loginTitle := "Login"
	loginLink := LoginURL
	user := ""

	acc, ok := auth.AccountFromContext(r.Context())
	if ok {
		user = acc.ID
		loginTitle = "Logout"
		loginLink = "/logout"
	}

	templateData := map[string]interface{}{
		"LoginTitle": loginTitle,
		"LoginURL":   loginLink,
		"Results":    data,
		"User":       user,
	}

	// add extra values
	for _, val := range vals {
		templateData[val.Key] = val.Value
	}

	if err := t.ExecuteTemplate(w, "layout",
		templateData,
	); err != nil {
		http.Error(w, "Error occurred:"+err.Error(), 500)
	}
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

type sortedServices struct {
	services []*registry.Service
}

func (s sortedServices) Len() int {
	return len(s.services)
}

func (s sortedServices) Less(i, j int) bool {
	return s.services[i].Name < s.services[j].Name
}

func (s sortedServices) Swap(i, j int) {
	s.services[i], s.services[j] = s.services[j], s.services[i]
}
