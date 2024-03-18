package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/itmosha/vk-internship-2024/pkg/jwtfuncs"
)

var (
	ErrAccessTokenNotProvided = errors.New("access token not provided")
	ErrInvalidAccessToken     = errors.New("invalid access token")
	ErrAccessTokenExpired     = errors.New("access token expired")
	ErrNotEnoughPermissions   = errors.New("not enough permissions")
)

// AuthMiddleware is a middleware to check authorization.
func AuthMiddleware(isAdminRequired bool, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		accessToken := extractTokenFromHeader(req.Header.Get("Authorization"))
		if accessToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrAccessTokenNotProvided.Error()})
			return
		}

		claims, isExpired, err := jwtfuncs.ExtractAccessTokenClaims(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrInvalidAccessToken.Error()})
			return
		}
		if isExpired {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": ErrAccessTokenExpired.Error()})
			return
		}
		if isAdminRequired && !claims.IsAdmin {
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
