package usecase

import (
	"context"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
)

// ActorRepo interface.
type ActorRepoInterface interface {
	Insert(ctx *context.Context, receivedActor *entity.Actor) (createdActor *entity.Actor, err error)
	Update(ctx *context.Context, id int, fields map[string]interface{}) (err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context) (actors []*entity.Actor, err error)
}

type ActorUsecase struct {
	actorRepo     ActorRepoInterface
	filmActorRepo FilmsActorsRepoInterface
}

// Create new ActorUsecase.
func NewActorUsecase(actorRepo ActorRepoInterface, filmActorRepo FilmsActorsRepoInterface) *ActorUsecase {
	return &ActorUsecase{
		actorRepo:     actorRepo,
		filmActorRepo: filmActorRepo,
	}
}

// Create a new actor.
func (uc *ActorUsecase) Create(ctx *context.Context, body *entity.ActorCreateBody) (actor *entity.Actor, err error) {
	actorToCreate := &entity.Actor{
		Name:      body.Name,
		Gender:    *body.Gender,
		BirthDate: body.BirthDate,
	}
	actor, err = uc.actorRepo.Insert(ctx, actorToCreate)
	return
}

// Update an actor by id.
func (uc *ActorUsecase) Update(ctx *context.Context, id int, body *entity.ActorUpdateBody) (err error) {
	fields := map[string]interface{}{}
	if len(body.Name) > 0 {
		fields["name"] = body.Name
	}
	if body.Gender != nil {
		fields["gender"] = *body.Gender
	}
	if len(body.BirthDate) > 0 {
		fields["birth_date"] = body.BirthDate
	}
	if len(fields) == 0 {
		return
	}
	err = uc.actorRepo.Update(ctx, id, fields)
	return
}

// Replace an actor by id.
func (uc *ActorUsecase) Replace(ctx *context.Context, id int, body *entity.ActorReplaceBody) (err error) {
	fields := map[string]interface{}{
		"name":       body.Name,
		"gender":     *body.Gender,
		"birth_date": body.BirthDate,
	}
	err = uc.actorRepo.Update(ctx, id, fields)
	return
}

// Delete an actor by id.
func (uc *ActorUsecase) Delete(ctx *context.Context, id int) (err error) {
	err = uc.actorRepo.Delete(ctx, id)
	return
}

// Get all actors.
func (uc *ActorUsecase) GetAll(ctx *context.Context) (actors []*entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}
