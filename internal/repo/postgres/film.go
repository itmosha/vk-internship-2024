package repo

import (
	"context"
	"database/sql"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	"github.com/itmosha/vk-internship-2024/pkg/postgres"
)

type FilmRepoPostgres struct {
	store *postgres.Postgres
}

// Create new FilmRepoPostgres.
func NewFilmRepoPostgres(store *postgres.Postgres) *FilmRepoPostgres {
	return &FilmRepoPostgres{store}
}

// Start a new transaction.
func (r *FilmRepoPostgres) NewTransaction(ctx *context.Context) (tx *sql.Tx, err error) {
	return r.store.DB.BeginTx(*ctx, nil)
}

// Insert a new Film with provided fields.
func (r *FilmRepoPostgres) Insert(ctx *context.Context, receivedFilm *entity.Film) (createdFilm *entity.Film, err error) {
	createdFilm = &entity.Film{}
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		INSERT INTO film (title, description, release_date, rating)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, description, release_date, rating;
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(*ctx, receivedFilm.Title, receivedFilm.Description, receivedFilm.ReleaseDate, receivedFilm.Rating).
		Scan(&createdFilm.ID, &createdFilm.Title, &createdFilm.Description, &createdFilm.ReleaseDate, &createdFilm.Rating)
	return
}

// Update provided fields of a Film by id.
func (r *FilmRepoPostgres) Update(ctx *context.Context, id int, fields map[string]interface{}) (updatedFilm *entity.Film, err error) {

	log.Panicln("not implemented")
	return
}

// Delete a Film by id.
func (r *FilmRepoPostgres) Delete(ctx *context.Context, id int) (err error) {

	log.Panicln("not implemented")
	return
}

// Get all Films.
func (r *FilmRepoPostgres) GetAll(ctx *context.Context, sortParams *entity.FilmSortParams, searchFields map[string]interface{}) (films []*entity.Film, err error) {

	log.Panicln("not implemented")
	return
}
