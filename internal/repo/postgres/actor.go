package repo

import (
	"context"
	"log"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	"github.com/itmosha/vk-internship-2024/pkg/postgres"
)

type ActorRepoPostgres struct {
	store *postgres.Postgres
}

// Create new ActorRepoPostgres.
func NewActorRepoPostgres(store *postgres.Postgres) *ActorRepoPostgres {
	return &ActorRepoPostgres{store}
}

// Insert a new Actor with provided fields.
func (r *ActorRepoPostgres) Insert(ctx *context.Context, receivedActor *entity.Actor) (createdActor *entity.Actor, err error) {
	createdActor = &entity.Actor{}
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		INSERT INTO actor (name, gender, birth_date)
		VALUES ($1, $2, $3)
		RETURNING id, name, gender, birth_date;
	`)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(*ctx, receivedActor.Name, receivedActor.Gender, receivedActor.BirthDate).
		Scan(&createdActor.ID, &createdActor.Name, &createdActor.Gender, &createdActor.BirthDate)
	return
}

// Update provided fields of a Actor by id.
func (r *ActorRepoPostgres) Update(context *context.Context, id int, fields map[string]interface{}) (updatedActor *entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}

// Delete a Actor by id.
func (r *ActorRepoPostgres) Delete(ctx *context.Context, id int) (err error) {

	log.Panicln("not implemented")
	return
}

// Get all Actors.
func (r *ActorRepoPostgres) GetAll(ctx *context.Context) (actors []*entity.Actor, err error) {

	log.Panicln("not implemented")
	return
}