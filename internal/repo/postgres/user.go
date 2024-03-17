package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	"github.com/itmosha/vk-internship-2024/pkg/postgres"
	"github.com/lib/pq"
)

type UserRepoPostgres struct {
	store *postgres.Postgres
}

func NewUserRepoPostgres(store *postgres.Postgres) *UserRepoPostgres {
	return &UserRepoPostgres{store}
}

func (r *UserRepoPostgres) Insert(ctx *context.Context, receivedUser *entity.User) (createdUser *entity.User, err error) {
	createdUser = &entity.User{}
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		INSERT INTO users (username)
		VALUES ($1)
		RETURNING id, username, is_admin;`)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(*ctx, receivedUser.Username).
		Scan(&createdUser.ID, &createdUser.Username, &createdUser.IsAdmin)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			err = ErrNonUniqueUsername
		}
		return
	}
	return
}

func (r *UserRepoPostgres) SelectByUsername(ctx *context.Context, username string) (user *entity.User, err error) {
	user = &entity.User{}
	stmt, err := r.store.DB.PrepareContext(*ctx, `
		SELECT id, username, is_admin
		FROM users
		WHERE username = $1;`)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(*ctx, username).
		Scan(&user.ID, &user.Username, &user.IsAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrUserNotFound
		}
	}
	return
}
