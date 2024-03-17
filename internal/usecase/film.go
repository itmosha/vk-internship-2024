package usecase

import (
	"context"
	"database/sql"
	"errors"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
)

// FilmUsecase interface.
type FilmRepoInterface interface {
	NewTransaction(ctx *context.Context) (tx *sql.Tx, err error)
	Insert(ctx *context.Context, receivedFilm *entity.Film) (createdFilm *entity.Film, err error)
	Update(ctx *context.Context, id int, fields map[string]interface{}) (err error)
	Delete(ctx *context.Context, id int) (err error)
	GetAllWithActors(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields *entity.FilmSearchParams) (films []*entity.FilmWithActors, err error)
}

// FilmActorRepo interface.
type FilmsActorsRepoInterface interface {
	Insert(ctx *context.Context, receivedFilmActor *entity.FilmActor) (createdFilmActor *entity.FilmActor, err error)
	Delete(ctx *context.Context, filmID, actorID int) (err error)
	SelectByFilmID(ctx *context.Context, filmID int) (filmsActors []*entity.FilmActor, err error)
	SelectByActorID(ctx *context.Context, actorID int) (filmsActors []*entity.FilmActor, err error)
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
		Rating:      *body.Rating,
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
func (uc *FilmUsecase) Update(ctx *context.Context, id int, body *entity.FilmUpdateBody) (err error) {
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

	fields := map[string]interface{}{}
	if len(body.Title) > 0 {
		fields["title"] = body.Title
	}
	if len(body.Description) > 0 {
		fields["description"] = body.Description
	}
	if len(body.ReleaseDate) > 0 {
		fields["release_date"] = body.ReleaseDate
	}
	if body.Rating != nil {
		fields["rating"] = *body.Rating
	}
	if len(fields) == 0 {
		return
	}
	err = uc.filmRepo.Update(ctx, id, fields)
	if err != nil {
		return
	}
	if body.ActorsIDs != nil {
		filmsActors, err_ := uc.filmsActorsRepo.SelectByFilmID(ctx, id)
		if err_ != nil {
			err = err_
			return
		}
		for _, filmActor := range filmsActors {
			err = uc.filmsActorsRepo.Delete(ctx, id, filmActor.ActorID)
			if err != nil {
				if errors.Is(err, repo.ErrFilmActorNotFound) {
					err = nil
				} else {
					return
				}
			}
		}
		for _, actorID := range body.ActorsIDs {
			_, err = uc.filmsActorsRepo.Insert(ctx, &entity.FilmActor{
				FilmID:  id,
				ActorID: actorID,
			})
			if err != nil {
				return
			}
		}
	}
	return
}

// Replace a film by id.
func (uc *FilmUsecase) Replace(ctx *context.Context, id int, body *entity.FilmReplaceBody) (err error) {
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
	fields := map[string]interface{}{
		"title":        body.Title,
		"description":  body.Description,
		"release_date": body.ReleaseDate,
		"rating":       *body.Rating,
	}

	err = uc.filmRepo.Update(ctx, id, fields)
	if err != nil {
		return
	}

	filmsActors, err_ := uc.filmsActorsRepo.SelectByFilmID(ctx, id)
	if err_ != nil {
		err = err_
		return
	}
	for _, filmActor := range filmsActors {
		err = uc.filmsActorsRepo.Delete(ctx, id, filmActor.ActorID)
		if err != nil {
			if errors.Is(err, repo.ErrFilmActorNotFound) {
				err = nil
			} else {
				return
			}
		}
	}
	for _, actorID := range body.ActorsIDs {
		_, err = uc.filmsActorsRepo.Insert(ctx, &entity.FilmActor{
			FilmID:  id,
			ActorID: actorID,
		})
		if err != nil {
			return
		}
	}
	return
}

// Delete a film by id.
func (uc *FilmUsecase) Delete(ctx *context.Context, id int) (err error) {
	err = uc.filmRepo.Delete(ctx, id)
	return
}

// Get all films.
func (uc *FilmUsecase) GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields *entity.FilmSearchParams) (films []*entity.FilmWithActors, err error) {
	films, err = uc.filmRepo.GetAllWithActors(ctx, sortParams, searchFields)
	return
}
