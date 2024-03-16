package usecase

import (
	"context"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
)

// ActorRepo interface.
type ActorRepoInterface interface {
	Insert(ctx *context.Context, receivedActor *entity.Actor) (createdActor *entity.Actor, err error)
	Update(ctx *context.Context, id int, fields map[string]interface{}) (updatedActor *entity.Actor, err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context) (actors []*entity.Actor, err error)
}

type ActorUsecase struct {
	actorRepo     ActorRepoInterface
	filmActorRepo FilmActorRepoInterface
}

// Create new ActorUsecase.
func NewActorUsecase(actorRepo ActorRepoInterface, filmActorRepo FilmActorRepoInterface) *ActorUsecase {
	return &ActorUsecase{
		actorRepo:     actorRepo,
		filmActorRepo: filmActorRepo,
	}
}

// Create a new actor.
func (uc *ActorUsecase) Create(ctx *context.Context, body *entity.ActorCreateBody) (actor *entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}

// Update an actor by id.
func (uc *ActorUsecase) Update(ctx *context.Context, id int, body *entity.ActorUpdateBody) (actor *entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}

// Replace an actor by id.
func (uc *ActorUsecase) Replace(ctx *context.Context, id int, body *entity.ActorReplaceBody) (actor *entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}

// Delete an actor by id.
func (uc *ActorUsecase) Delete(ctx *context.Context, id int) (err error) {

	log.Panicln("not implemented")
	return
}

// Get all actors.
func (uc *ActorUsecase) GetAll(ctx *context.Context) (actors []*entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}
