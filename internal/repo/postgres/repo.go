package repo

import "errors"

var (
	ErrFilmNotFound      = errors.New("film with provided id was not found")
	ErrActorNotFound     = errors.New("actor with provided id was not found")
	ErrFilmActorNotFound = errors.New("film_actor with provided film_id and actor_id was not found")
)
