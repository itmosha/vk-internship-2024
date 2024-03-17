package http_server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	jwtfuncs "github.com/itmosha/vk-internship-2024/pkg/jwt_funcs"
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
	GetAllWithFilms() http.HandlerFunc
}

// User handler interface.
type UserHandlerInterface interface {
	Register() http.HandlerFunc
	Login() http.HandlerFunc
}

var (
	ErrAccessTokenNotProvided = errors.New("access token not provided")
	ErrInvalidAccessToken     = errors.New("invalid access token")
	ErrAccessTokenExpired     = errors.New("access token expired")
	ErrNotEnoughPermissions   = errors.New("not enough permissions")
)

// Router struct.
type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

// Create new Router.
func NewRouter(filmHandler FilmHandlerInterface, actorHandler ActorHandlerInterface, userHandler UserHandlerInterface) (router *Router) {
	router = &Router{routes: make(map[string]map[string]http.HandlerFunc)}

	// Ping endpoint
	router.HandleFunc("/ping", http.MethodGet, func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("pong"))
	})

	// Film endpoints
	router.HandleFunc("/api/films/", http.MethodPost, AuthMiddleware(true, filmHandler.Create()))
	router.HandleFunc("/api/films/{id}/", http.MethodPatch, AuthMiddleware(true, filmHandler.Update()))
	router.HandleFunc("/api/films/{id}/", http.MethodPut, AuthMiddleware(true, filmHandler.Replace()))
	router.HandleFunc("/api/films/{id}", http.MethodDelete, AuthMiddleware(true, filmHandler.Delete()))
	router.HandleFunc("/api/films", http.MethodGet, AuthMiddleware(false, filmHandler.GetAll()))

	// Actor endpoints
	router.HandleFunc("/api/actors/", http.MethodPost, AuthMiddleware(true, actorHandler.Create()))
	router.HandleFunc("/api/actors/{id}/", http.MethodPatch, AuthMiddleware(true, actorHandler.Update()))
	router.HandleFunc("/api/actors/{id}/", http.MethodPut, AuthMiddleware(true, actorHandler.Replace()))
	router.HandleFunc("/api/actors/{id}", http.MethodDelete, AuthMiddleware(true, actorHandler.Delete()))
	router.HandleFunc("/api/actors", http.MethodGet, AuthMiddleware(false, actorHandler.GetAllWithFilms()))

	// User endpoints
	router.HandleFunc("/api/auth/register/", http.MethodPost, userHandler.Register())
	router.HandleFunc("/api/auth/login/", http.MethodPost, userHandler.Login())
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

// AuthMiddleware is a middleware to check authorization.
func AuthMiddleware(isAdminRequred bool, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		accessToken := extractTokenFromHeader(req.Header.Get("Authorization"))
		if accessToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrAccessTokenNotProvided.Error()})
			return
		}

		claims, isExpired, err := jwtfuncs.ExtractAccessTokenClaims(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrInvalidAccessToken.Error()})
			return
		}
		if isExpired {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrAccessTokenExpired.Error()})
			return
		}
		if isAdminRequred && !claims.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrNotEnoughPermissions.Error()})
			return
		}
		next.ServeHTTP(w, req)
	}
}

// Extract token from "Authorization" header.
func extractTokenFromHeader(header string) string {
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}

// Add a value to request context.
func withValue(req *http.Request, key, val string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, key, val)
	return req.WithContext(ctx)
}
