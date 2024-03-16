package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
)

type FilmUsecaseInterface interface {
	Create(ctx *context.Context, body *entity.FilmCreateBody) (film *entity.Film, err error)
	Update(ctx *context.Context, id int, body *entity.FilmUpdateBody) (err error)
	Replace(ctx *context.Context, id int, body *entity.FilmReplaceBody) (err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields map[string]interface{}) (films []*entity.Film, err error)
}

type FilmHandler struct {
	filmUsecase FilmUsecaseInterface
}

// Create new FilmHandler.
func NewFilmHander(filmUsecase FilmUsecaseInterface) *FilmHandler {
	return &FilmHandler{filmUsecase}
}

// Create a new film.
func (h *FilmHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.FilmCreateBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateFilmCreateBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		film, err := h.filmUsecase.Create(&ctx, body)
		if err != nil {
			switch err {
			case repo.ErrActorNotFound:
				returnError(w, http.StatusBadRequest, err)
				return
			default:
				returnError(w, http.StatusInternalServerError, ErrServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(film)
	}
}

// Update a film by id.
func (h *FilmHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		id, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		body, err := readBodyToStruct(r, &entity.FilmUpdateBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateFilmUpdateBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.filmUsecase.Update(&ctx, id, body)
		if err != nil {
			fmt.Println(err)
			switch err {
			case repo.ErrFilmNotFound:
				returnError(w, http.StatusBadRequest, err)
				return
			default:
				returnError(w, http.StatusInternalServerError, ErrServerError)
				return
			}
		}
	}
}

// Replace a film by id.
func (h *FilmHandler) Replace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Delete a film by id.
func (h *FilmHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Get all films.
func (h *FilmHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
