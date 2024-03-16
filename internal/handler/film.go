package handler

import (
	"context"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
)

type FilmUsecaseInterface interface {
	Create(ctx *context.Context, body *entity.FilmCreateBody) (film *entity.Film, err error)
	Update(ctx *context.Context, id int, body *entity.FilmUpdateBody) (film *entity.Film, err error)
	Replace(ctx *context.Context, id int, body *entity.FilmReplaceBody) (film *entity.Film, err error)
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
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Update a film by id.
func (h *FilmHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
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
