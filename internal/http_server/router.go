package http_server

import "net/http"

type Router struct {
	routes map[string]http.HandlerFunc
}

// Add new route to router.
func (r *Router) HandleFunc(path string, handler http.HandlerFunc) {
	r.routes[path] = handler
}

// Implement http.Handler interface's ServeHTTP method.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := r.routes[req.URL.Path]; ok {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

// Create new router.
func NewRouter() (router *Router) {
	router = &Router{routes: make(map[string]http.HandlerFunc)}

	router.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	})
	return
}
