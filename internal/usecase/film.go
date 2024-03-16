package usecase

import (
	"context"
	"database/sql"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
)

// FilmUsecase interface.
type FilmRepoInterface interface {
	NewTransaction(ctx *context.Context) (tx *sql.Tx, err error)
	Insert(ctx *context.Context, receivedFilm *entity.Film) (createdFilm *entity.Film, err error)
	Update(ctx *context.Context, id int, fields map[string]interface{}) (updatedFilm *entity.Film, err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields map[string]interface{}) (films []*entity.Film, err error)
}

// FilmActorRepo interface.
type FilmsActorsRepoInterface interface {
	Insert(ctx *context.Context, receivedFilmActor *entity.FilmActor) (createdFilmActor *entity.FilmActor, err error)
	Delete(ctx *context.Context, filmID, actorID int) (err error)
}

type FilmUsecase struct {
	filmRepo        FilmRepoInterface
	actorRepo       ActorRepoInterface
	filmsActorsRepo FilmsActorsRepoInterface
}

// Create new FilmUsecase.
func NewFilmUsecase(filmRepo FilmRepoInterface, actorRepo ActorRepoInterface, filmsActorsRepo FilmsActorsRepoInterface) *FilmUsecase {
	return &FilmUsecase{
		filmRepo:        filmRepo,
		actorRepo:       actorRepo,
		filmsActorsRepo: filmsActorsRepo,
	}
}

// Create a new film.
func (uc *FilmUsecase) Create(ctx *context.Context, body *entity.FilmCreateBody) (film *entity.Film, err error) {
	tx, err := uc.filmRepo.NewTransaction(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	filmToCreate := &entity.Film{
		Title:       body.Title,
		Description: body.Description,
		ReleaseDate: body.ReleaseDate,
		Rating:      body.Rating,
	}
	film, err = uc.filmRepo.Insert(ctx, filmToCreate)
	if err != nil {
		return
	}
	for _, actorID := range body.ActorsIDs {
		_, err = uc.filmsActorsRepo.Insert(ctx, &entity.FilmActor{
			FilmID:  film.ID,
			ActorID: actorID,
		})
		if err != nil {
			return
		}
	}
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
