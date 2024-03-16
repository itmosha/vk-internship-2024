package http_server

import (
	"context"
	"net/http"
	"strings"
)

// Film handler interface.
type FilmHandlerInterface interface {
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Replace() http.HandlerFunc
	Delete() http.HandlerFunc
	GetAll() http.HandlerFunc
}

// Actor handler interface.
type ActorHandlerInterface interface {
	Create() http.HandlerFunc
	Update() http.HandlerFunc
	Replace() http.HandlerFunc
	Delete() http.HandlerFunc
	GetAll() http.HandlerFunc
}

// Router struct.
type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

// Create new Router.
func NewRouter(filmHandler FilmHandlerInterface, actorHandler ActorHandlerInterface) (router *Router) {
	router = &Router{routes: make(map[string]map[string]http.HandlerFunc)}

	// Ping endpoint
	router.HandleFunc("/ping", http.MethodGet, func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	})

	// Film endpoints
	router.HandleFunc("/api/films/", http.MethodPost, filmHandler.Create())
	router.HandleFunc("/api/films/{id}/", http.MethodPatch, filmHandler.Update())
	router.HandleFunc("/api/films/{id}/", http.MethodPut, filmHandler.Replace())
	router.HandleFunc("/api/films/{id}", http.MethodDelete, filmHandler.Delete())
	router.HandleFunc("/api/films", http.MethodGet, filmHandler.GetAll())

	// Actor endpoints
	router.HandleFunc("/api/actors/", http.MethodPost, actorHandler.Create())
	router.HandleFunc("/api/actors/{id}/", http.MethodPatch, actorHandler.Update())
	router.HandleFunc("/api/actors/{id}/", http.MethodPut, actorHandler.Replace())
	router.HandleFunc("/api/actors/{id}", http.MethodDelete, actorHandler.Delete())
	router.HandleFunc("/api/actors", http.MethodGet, actorHandler.GetAll())
	return
}

// Register handler.
func (r *Router) HandleFunc(path string, method string, handler http.HandlerFunc) {
	if _, ok := r.routes[path]; !ok {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}

// ServeHTTP.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	var handler http.HandlerFunc
	for registeredPath, methodHandlers := range r.routes {
		if matched, params := pathMatch(registeredPath, path); matched {
			if h, ok := methodHandlers[method]; ok {
				handler = h
				for key, value := range params {
					req = withValue(req, key, value)
				}
				break
			}
		}
	}
	if handler != nil {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

// Match request path with registered path.
func pathMatch(registeredPath, reqPath string) (bool, map[string]string) {
	registeredParts := strings.Split(registeredPath, "/")
	reqParts := strings.Split(reqPath, "/")
	if len(registeredParts) != len(reqParts) {
		return false, nil
	}
	params := make(map[string]string)
	for i, part := range registeredParts {
		if part != reqParts[i] {
			if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
				params[strings.Trim(part, "{}")] = reqParts[i]
			} else {
				return false, nil
			}
		}
	}
	return true, params
}

// Add a value to request context.
func withValue(req *http.Request, key, val string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, key, val)
	return req.WithContext(ctx)
}
