package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
	"github.com/itmosha/vk-internship-2024/pkg/logger"
)

type FilmUsecaseInterface interface {
	Create(ctx *context.Context, body *entity.FilmCreateBody) (film *entity.Film, err error)
	Update(ctx *context.Context, id int, body *entity.FilmUpdateBody) (err error)
	Replace(ctx *context.Context, id int, body *entity.FilmReplaceBody) (err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields *entity.FilmSearchParams) (films []*entity.FilmWithActors, err error)
}

type FilmHandler struct {
	filmUsecase FilmUsecaseInterface
	logger      *logger.Logger
}

// Create new FilmHandler.
func NewFilmHander(filmUsecase FilmUsecaseInterface, logger *logger.Logger) *FilmHandler {
	return &FilmHandler{filmUsecase, logger}
}

// @Title Create film
// @Description Create a new film.
// @Param body body entity.FilmCreateBody true "Create film body"
// @Success 201 {object} entity.Film
// @Failure 400 {object} RequestError
// @Failure 401 {object} RequestError
// @Failure 403 {object} RequestError
// @Failure 500 {object} RequestError
// @Resource Films
// @Route /api/films/ [post]
func (h *FilmHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			h.logger.Log(r, http.StatusBadRequest, ErrEmptyBody)
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.FilmCreateBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateFilmCreateBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		film, err := h.filmUsecase.Create(&ctx, body)
		if err != nil {
			switch err {
			case repo.ErrActorNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(film)
		h.logger.Log(r, http.StatusCreated, nil)
	}
}

// @Title Update film
// @Description Update a film by id.
// @Param id path integer true "Film ID"
// @Param body body entity.FilmUpdateBody true "Update film body"
// @Success 200 {}
// @Failure 400 {object} RequestError
// @Failure 401 {object} RequestError
// @Failure 403 {object} RequestError
// @Failure 500 {object} RequestError
// @Resource Films
// @Route /api/films/{id}/ [patch]
func (h *FilmHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			h.logger.Log(r, http.StatusBadRequest, ErrEmptyBody)
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		id, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		body, err := readBodyToStruct(r, &entity.FilmUpdateBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateFilmUpdateBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.filmUsecase.Update(&ctx, id, body)
		if err != nil {
			fmt.Println(err)
			switch err {
			case repo.ErrFilmNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			case repo.ErrActorNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		h.logger.Log(r, http.StatusOK, nil)
	}
}

// @Title Replace film
// @Description Replace a film by id.
// @Param id path integer true "Film ID"
// @Param body body entity.FilmReplaceBody true "Replace film body"
// @Success 200 {}
// @Failure 400 {object} RequestError
// @Failure 401 {object} RequestError
// @Failure 403 {object} RequestError
// @Failure 500 {object} RequestError
// @Resource Films
// @Route /api/films/{id}/ [put]
func (h *FilmHandler) Replace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			h.logger.Log(r, http.StatusBadRequest, ErrEmptyBody)
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		id, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		body, err := readBodyToStruct(r, &entity.FilmReplaceBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateFilmReplaceBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.filmUsecase.Replace(&ctx, id, body)
		if err != nil {
			switch err {
			case repo.ErrFilmNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			case repo.ErrActorNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		h.logger.Log(r, http.StatusOK, nil)
	}
}

// @Title Delete film
// @Description Delete a film by id.
// @Param id path integer true "Film ID"
// @Success 204 {}
// @Failure 400 {object} RequestError
// @Failure 401 {object} RequestError
// @Failure 403 {object} RequestError
// @Failure 500 {object} RequestError
// @Resource Films
// @Route /api/films/{id} [delete]
func (h *FilmHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		ctx := context.Background()
		err = h.filmUsecase.Delete(&ctx, id)
		if err != nil {
			switch err {
			case repo.ErrFilmNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		w.WriteHeader(http.StatusNoContent)
		h.logger.Log(r, http.StatusNoContent, nil)
	}
}

// @Title Get all films
// @Description Get all films with seaching and sorting params.
// @Param sort_by query string true "Sort by field" Enums(title,rating,release_date)
// @Param sort_order query string true "Sort order" Enums(asc,desc)
// @Param title query string true "Search by title"
// @Param actor_name query string true "Search by actor name"
// @Success 200 {array} entity.FilmWithActors
// @Failure 401 {object} RequestError
// @Failure 500 {object} RequestError
// @Resource Films
// @Route /api/films [get]
func (h *FilmHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Refactor this method
		query := r.URL.Query()
		sortField := query.Get("sort_by")
		if sortField == "" {
			sortField = entity.FilmDefaultSortField
		} else if !entity.IsValidParam("sort_field", sortField) {
			h.logger.Log(r, http.StatusBadRequest, ErrInvalidSortByParam)
			returnError(w, http.StatusBadRequest, ErrInvalidSortByParam)
			return
		}
		sortOrder := query.Get("order")
		if sortOrder == "" {
			sortOrder = entity.FilmDefaultSortOrder
		} else if !entity.IsValidParam("sort_order", sortOrder) {
			h.logger.Log(r, http.StatusBadRequest, ErrInvalidOrderParam)
			returnError(w, http.StatusBadRequest, ErrInvalidOrderParam)
			return
		}
		sortParams := &entity.FilmSortParams{
			Field: sortField,
			Order: sortOrder,
		}
		searchParams := &entity.FilmSearchParams{
			Title:     query.Get("title"),
			ActorName: query.Get("actor_name"),
		}
		ctx := context.Background()
		films, err := h.filmUsecase.GetAll(&ctx, sortParams, searchParams)
		if err != nil {
			fmt.Println(err)
			switch err {
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		if films == nil { // Handle empty films slice, return [] instead of nil
			json.NewEncoder(w).Encode([]struct{}{})
		} else {
			json.NewEncoder(w).Encode(films)
		}
		h.logger.Log(r, http.StatusOK, nil)
	}
}
