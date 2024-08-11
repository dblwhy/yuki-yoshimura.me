package http

import "net/http"

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, exists := r.routes[req.Method]; exists {
		if handler, exists := handlers[req.URL.Path]; exists {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
