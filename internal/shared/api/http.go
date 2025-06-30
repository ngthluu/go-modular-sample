package api

import (
	"go-modular-sample/internal/shared/constant"
	"net/http"
)

type Controller interface {
	Path() string
	Prepare(http.ResponseWriter, *http.Request) bool
	Perform(http.ResponseWriter, *http.Request)
}

func NewServeMux(routes []Controller) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Path(), _HttpControllerWrapper{
			c: route,
		})
	}
	return mux
}

func NewHTTPServer(mux *http.ServeMux) *http.Server {
	return &http.Server{Addr: ":" + constant.EnvPort, Handler: mux}
}

// HTTP controller wrapper
type _HttpControllerWrapper struct {
	http.Handler
	c Controller
}

func (wrapper _HttpControllerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if wrapper.c.Prepare(w, r) {
		wrapper.c.Perform(w, r)
	}
}
