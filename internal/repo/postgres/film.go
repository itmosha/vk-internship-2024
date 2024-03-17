package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

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
func (r *FilmRepoPostgres) Update(ctx *context.Context, id int, fields map[string]interface{}) (err error) {
	// TODO: Use prepared statements
	query := `UPDATE film SET id = id, `
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
		err = ErrFilmNotFound
	}
	return
}

// Delete a Film by id.
func (r *FilmRepoPostgres) Delete(ctx *context.Context, id int) (err error) {
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		DELETE FROM film
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
		err = ErrFilmNotFound
	}
	return
}

// Get all films.
func (r *FilmRepoPostgres) GetAllWithActors(ctx *context.Context, sortParams *entity.FilmSortParams, searchParams *entity.FilmSearchParams) (films []*entity.FilmWithActors, err error) {
	// TODO: Refactor this mess
	query := `
		SELECT f.id, f.title, f.description, f.release_date, f.rating, ARRAY_AGG(a.id) AS actor_ids
        FROM film f
        LEFT JOIN films_actors fa ON f.id = fa.film_id
        LEFT JOIN actor a ON fa.actor_id = a.id`
	var conditions []string
	var args []interface{}
	paramIndex := 1
	if searchParams != nil {
		if searchParams.Title != "" {
			conditions = append(conditions, "f.title ILIKE $"+strconv.Itoa(paramIndex))
			args = append(args, "%"+searchParams.Title+"%")
			paramIndex++
		}
		if searchParams.ActorName != "" {
			conditions = append(conditions, "a.name ILIKE $"+strconv.Itoa(paramIndex))
			args = append(args, "%"+searchParams.ActorName+"%")
			paramIndex++
		}
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	query += " GROUP BY f.id"
	if sortParams != nil {
		query += " ORDER BY " + sortParams.Field + " " + sortParams.Order
	}
	rows, err := r.store.DB.QueryContext(*ctx, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		film := entity.FilmWithActors{}
		var actorsIDsRaw []byte
		if err = rows.Scan(&film.ID, &film.Title, &film.Description, &film.ReleaseDate, &film.Rating, &actorsIDsRaw); err != nil {
			return
		}
		// TODO: Refactor this also
		actorsIDsStr := string(actorsIDsRaw)
		if actorsIDsStr == "{NULL}" {
			film.ActorsIDs = []int{}
		} else {
			actorsIDsStr = strings.Trim(actorsIDsStr, "{}")
			actorsIDsStrSlice := strings.Split(actorsIDsStr, ",")
			actorsIDs := make([]int, len(actorsIDsStrSlice))
			for i, idStr := range actorsIDsStrSlice {
				actorsIDs[i], _ = strconv.Atoi(idStr)
			}
			film.ActorsIDs = actorsIDs
		}
		films = append(films, &film)
	}
	return
}
