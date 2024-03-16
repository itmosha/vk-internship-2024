package repo

import (
	"context"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	"github.com/itmosha/vk-internship-2024/pkg/postgres"
)

type FilmActorRepoPostgres struct {
	store *postgres.Postgres
}

// Create new FilmActorRepoPostgres.
func NewFilmActorRepoPostgres(store *postgres.Postgres) *FilmActorRepoPostgres {
	return &FilmActorRepoPostgres{store}
}

// Insert a new FilmActor with provided fields.
func (r *FilmActorRepoPostgres) Insert(ctx *context.Context, receivedFilmActor *entity.FilmActor) (createdFilmActor *entity.FilmActor, err error) {

	log.Panicln("not implemented")
	return
}

// Delete a FilmActor by film_id and actor_id.
func (r *FilmActorRepoPostgres) Delete(ctx *context.Context, filmID, actorID int) (err error) {

	log.Panicln("not implemented")
	return
}
