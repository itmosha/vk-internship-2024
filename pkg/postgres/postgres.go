package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	*sql.DB
}

// Create a new Postgres connection.
func NewPostgres(address, user, password, name string) (pg *Postgres, err error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		user, password, address, name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	pg = &Postgres{db}
	return
}
