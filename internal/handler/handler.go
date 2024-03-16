package handler

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrDecodeBody = errors.New("could not decode request body")
	ErrEmptyBody  = errors.New("empty request body")

	ErrServerError = errors.New("internal server error")
)

func readBodyToStruct[T any](r *http.Request, out *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		return nil, ErrDecodeBody
	}
	return out, nil
}

func isEmptyBody(r *http.Request) bool {
	return r.ContentLength == 0
}

func returnError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
}
