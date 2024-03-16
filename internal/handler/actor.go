package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
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
}

func NewActorHandler(actorUsecase ActorUsecaseInterface) *ActorHandler {
	return &ActorHandler{actorUsecase}
}

// Create a new actor.
func (h *ActorHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.ActorCreateBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateActorCreateBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		actor, err := h.actorUsecase.Create(&ctx, body)
		if err != nil {
			returnError(w, http.StatusInternalServerError, ErrServerError)
			return
		}
		json.NewEncoder(w).Encode(actor)
	}
}

// Update an actor by id.
func (h *ActorHandler) Update() http.HandlerFunc {
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
		body, err := readBodyToStruct(r, &entity.ActorUpdateBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateActorUpdateBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.actorUsecase.Update(&ctx, id, body)
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
	}
}

// Replace an actor by id.
func (h *ActorHandler) Replace() http.HandlerFunc {
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
		body, err := readBodyToStruct(r, &entity.ActorReplaceBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateActorReplaceBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		err = h.actorUsecase.Replace(&ctx, id, body)
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
	}
}

// Delete an actor by id.
func (h *ActorHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := extractIDFromPath(r.URL.Path)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		ctx := context.Background()
		err = h.actorUsecase.Delete(&ctx, id)
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
		w.WriteHeader(http.StatusNoContent)
	}
}

// Get all actors.
func (h *ActorHandler) GetAllWithFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		actors, err := h.actorUsecase.GetAllWithFilms(&ctx)
		if err != nil {
			fmt.Println(err)
			returnError(w, http.StatusInternalServerError, ErrServerError)
			return
		}
		json.NewEncoder(w).Encode(actors)
	}
}
