package handler

import (
	"context"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
)

type ActorUsecaseInterface interface {
	Create(ctx *context.Context, body *entity.ActorCreateBody) (actor *entity.Actor, err error)
	Update(ctx *context.Context, id int, body *entity.ActorUpdateBody) (actor *entity.Actor, err error)
	Replace(ctx *context.Context, id int, body *entity.ActorReplaceBody) (actor *entity.Actor, err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context) (actors []*entity.Actor, err error)
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
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Update an actor by id.
func (h *ActorHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Replace an actor by id.
func (h *ActorHandler) Replace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Delete an actor by id.
func (h *ActorHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

// Get all actors.
func (h *ActorHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
