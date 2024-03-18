package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrInvalidPathParameter = errors.New("invalid path parameter")
	ErrDecodeBody           = errors.New("could not decode request body")
	ErrEmptyBody            = errors.New("empty request body")

	ErrInvalidSortByParam   = errors.New("invalid sort_by query parameter, should be one of: title, rating, release_date")
	ErrInvalidOrderParam    = errors.New("invalid order query parameter, should be one of: asc, desc")
	ErrInvalidSearchByParam = errors.New("invalid search_by query parameter, should be one of: title, actor_name")

	ErrServerError = errors.New("internal server error")
)

type RequestError struct {
	Message string `json:"message"`
}

func readBodyToStruct[T any](r *http.Request, out *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		return nil, ErrDecodeBody
	}
	return out, nil
}

func extractIDFromPath(path string) (id int, err error) {
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			id, err = strconv.Atoi(parts[i])
			if err != nil {
				err = ErrInvalidPathParameter
			}
			return
		}
	}
	return
}

func isEmptyBody(r *http.Request) bool {
	return r.ContentLength == 0
}

func returnError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(RequestError{Message: err.Error()})
}
