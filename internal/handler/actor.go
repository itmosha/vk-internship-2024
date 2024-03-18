package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
	"github.com/itmosha/vk-internship-2024/pkg/logger"
)

type ActorUsecaseInterface interface {
	Create(ctx *context.Context, body *entity.ActorCreateBody) (actor *entity.Actor, err error)
	Update(ctx *context.Context, id int, body *entity.ActorUpdateBody) (err error)
	Replace(ctx *context.Context, id int, body *entity.ActorReplaceBody) (err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAllWithFilms(ctx *context.Context) (actors []*entity.ActorWithFilms, err error)
}

type ActorHandler struct {
	actorUsecase ActorUsecaseInterface
	logger       *logger.Logger
}

func NewActorHandler(actorUsecase ActorUsecaseInterface, logger *logger.Logger) *ActorHandler {
	return &ActorHandler{actorUsecase, logger}
}

// Create a new actor.
func (h *ActorHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			h.logger.Log(r, http.StatusBadRequest, ErrEmptyBody)
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.ActorCreateBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateActorCreateBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		actor, err := h.actorUsecase.Create(&ctx, body)
		if err != nil {
			switch err {
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(actor)
		h.logger.Log(r, http.StatusCreated, nil)
	}
}

// Update an actor by id.
func (h *ActorHandler) Update() http.HandlerFunc {
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
		body, err := readBodyToStruct(r, &entity.ActorUpdateBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateActorUpdateBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.actorUsecase.Update(&ctx, id, body)
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
		h.logger.Log(r, http.StatusOK, nil)
	}
}

// Replace an actor by id.
func (h *ActorHandler) Replace() http.HandlerFunc {
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
		body, err := readBodyToStruct(r, &entity.ActorReplaceBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateActorReplaceBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.actorUsecase.Replace(&ctx, id, body)
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
		h.logger.Log(r, http.StatusOK, nil)
	}
}

// Delete an actor by id.
func (h *ActorHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		ctx := context.Background()
		err = h.actorUsecase.Delete(&ctx, id)
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
		w.WriteHeader(http.StatusNoContent)
		h.logger.Log(r, http.StatusNoContent, nil)
	}
}

// Get all actors.
func (h *ActorHandler) GetAllWithFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		actors, err := h.actorUsecase.GetAllWithFilms(&ctx)
		if err != nil {
			switch err {
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}

		if actors == nil { // Handle empty actors slice, return [] instead of nil
			json.NewEncoder(w).Encode([]struct{}{})
		} else {
			json.NewEncoder(w).Encode(actors)
		}
		h.logger.Log(r, http.StatusOK, nil)
	}
}
