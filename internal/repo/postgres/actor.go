package repo

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

// Update provided fields of an Actor by id.
func (r *ActorRepoPostgres) Update(ctx *context.Context, id int, fields map[string]interface{}) (err error) {
	// TODO: Use prepared statements
	query := `UPDATE actor SET id = id, `
	values := make([]interface{}, 0)
	idx := 1
	for field, value := range fields {
		query += field + "=$" + fmt.Sprint(idx) + ", "
		values = append(values, value)
		idx++
	}
	query = query[:len(query)-2] + " WHERE id=$" + fmt.Sprint(idx)
	values = append(values, id)

	res, err := r.store.DB.ExecContext(*ctx, query, values...)
	if err != nil {
		return
	}
	var cntRows int64
	cntRows, err = res.RowsAffected()
	if err != nil {
		return
	} else if cntRows == 0 {
		err = ErrActorNotFound
	}
	return
}

// Delete an Actor by id.
func (r *ActorRepoPostgres) Delete(ctx *context.Context, id int) (err error) {
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		DELETE FROM actor
		WHERE id = $1;`)
	if err != nil {
		return
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(*ctx, id)
	if err != nil {
		return
	}
	if cntRows, _ := res.RowsAffected(); cntRows == 0 {
		err = ErrActorNotFound
	}
	return
}

// Get all Actors.
func (r *ActorRepoPostgres) GetAllWithFilms(ctx *context.Context) (actors []*entity.ActorWithFilms, err error) {
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		SELECT a.id, a.name, a.gender, a.birth_date, ARRAY_AGG(f.id) AS films_ids
		FROM actor a
		LEFT JOIN films_actors fa ON a.id = fa.actor_id
		LEFT JOIN film f ON fa.film_id = f.id
		GROUP BY a.id;`)

	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(*ctx)
	if err != nil {
		return
	}
	for rows.Next() {
		actor := &entity.ActorWithFilms{}
		var filmsIDsRaw []byte
		err = rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDate, &filmsIDsRaw)
		if err != nil {
			return
		}
		// TODO: Refactor this
		filmsIDsStr := string(filmsIDsRaw)
		if filmsIDsStr == "{NULL}" {
			actor.FilmsIDs = []int{}
		} else {
			filmsIDsStr = strings.Trim(filmsIDsStr, "{}")
			filmsIDsStrSlice := strings.Split(filmsIDsStr, ",")
			filmsIDs := make([]int, len(filmsIDsStrSlice))
			for i, idStr := range filmsIDsStrSlice {
				filmsIDs[i], _ = strconv.Atoi(idStr)
			}
			actor.FilmsIDs = filmsIDs
		}
		actors = append(actors, actor)
	}
	return
}
