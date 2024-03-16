package repo

import (
	"context"
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

// Insert a new Film with provided fields.
func (r *FilmRepoPostgres) Insert(ctx *context.Context, receivedFilm *entity.Film) (createdFilm *entity.Film, err error) {

	log.Panicln("not implemented")
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
