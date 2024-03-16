package repo

import "errors"

var (
	ErrFilmNotFound  = errors.New("film with provided id was not found")
	ErrActorNotFound = errors.New("actor with provided id was not found")
)
