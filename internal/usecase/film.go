package usecase

import (
	"context"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
)

// FilmUsecase interface.
type FilmRepoInterface interface {
	Insert(ctx *context.Context, receivedFilm *entity.Film) (createdFilm *entity.Film, err error)
	Update(ctx *context.Context, id int, fields map[string]interface{}) (updatedFilm *entity.Film, err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields map[string]interface{}) (films []*entity.Film, err error)
}

// FilmActorRepo interface.
type FilmActorRepoInterface interface {
	Insert(ctx *context.Context, receivedFilmActor *entity.FilmActor) (createdFilmActor *entity.FilmActor, err error)
	Delete(ctx *context.Context, filmID, actorID int) (err error)
}

type FilmUsecase struct {
	filmRepo      FilmRepoInterface
	actorRepo     ActorRepoInterface
	filmActorRepo FilmActorRepoInterface
}

// Create new FilmUsecase.
func NewFilmUsecase(filmRepo FilmRepoInterface, actorRepo ActorRepoInterface, filmActorRepo FilmActorRepoInterface) *FilmUsecase {
	return &FilmUsecase{
		filmRepo:      filmRepo,
		actorRepo:     actorRepo,
		filmActorRepo: filmActorRepo,
	}
}

// Create a new film.
func (uc *FilmUsecase) Create(ctx *context.Context, body *entity.FilmCreateBody) (film *entity.Film, err error) {

	log.Panicln("not implemented")
	return
}

// Update a film by id.
func (uc *FilmUsecase) Update(ctx *context.Context, id int, body *entity.FilmUpdateBody) (film *entity.Film, err error) {

	log.Panicln("not implemented")
	return
}

// Replace a film by id.
func (uc *FilmUsecase) Replace(ctx *context.Context, id int, body *entity.FilmReplaceBody) (film *entity.Film, err error) {

	log.Panicln("not implemented")
	return
}

// Delete a film by id.
func (uc *FilmUsecase) Delete(ctx *context.Context, id int) (err error) {

	log.Panicln("not implemented")
	return
}

// Get all films.
func (uc *FilmUsecase) GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields map[string]interface{}) (films []*entity.Film, err error) {

	log.Panicln("not implemented")
	return
}
