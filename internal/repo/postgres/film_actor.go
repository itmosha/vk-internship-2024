package repo

import (
	"context"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	"github.com/itmosha/vk-internship-2024/pkg/postgres"
	"github.com/lib/pq"
)

type FilmsActorsRepoPostgres struct {
	store *postgres.Postgres
}

// Create new FilmActorRepoPostgres.
func NewFilmsActorsRepoPostgres(store *postgres.Postgres) *FilmsActorsRepoPostgres {
	return &FilmsActorsRepoPostgres{store}
}

// Insert a new FilmActor with provided fields.
func (r *FilmsActorsRepoPostgres) Insert(ctx *context.Context, receivedFilmActor *entity.FilmActor) (createdFilmActor *entity.FilmActor, err error) {
	createdFilmActor = &entity.FilmActor{}
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		INSERT INTO films_actors (film_id, actor_id)
		VALUES ($1, $2)
		RETURNING film_id, actor_id;`)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(*ctx, receivedFilmActor.FilmID, receivedFilmActor.ActorID).
		Scan(&createdFilmActor.FilmID, &createdFilmActor.ActorID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Constraint == "films_actors_actor_id_fkey" {
				// TODO: return id of actor that doesn't exist
				err = ErrActorNotFound
			} else if pqErr.Constraint == "films_actors_film_id_fkey" {
				// TODO: return id of film that doesn't exist
				err = ErrFilmNotFound
			}
		}
	}
	return
}

// Delete a FilmActor by film_id and actor_id.
func (r *FilmsActorsRepoPostgres) Delete(ctx *context.Context, filmID, actorID int) (err error) {

	log.Panicln("not implemented")
	return
}
