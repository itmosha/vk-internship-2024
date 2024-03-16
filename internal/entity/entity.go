package entity

import "errors"

var (
	ErrInvalidActorNameLength = errors.New("invalid length of name field, should be of length 1 to 100")
	ErrInvalidActorBirthDate  = errors.New("invalid format of birth_date field, must be in format 01.02.2006")

	ErrInvalidFilmTitleLength       = errors.New("invalid length of title field, should be of length 1 to 150")
	ErrInvalidFilmDescriptionLength = errors.New("invalid length of description field, should be of length 1 to 1000")
	ErrInvalidFilmReleaseDate       = errors.New("invalid format of release_date field, must be in format 01.02.2006")
	ErrInvalidFilmRating            = errors.New("invalid value of rating field, should be in range 0 to 10")
	ErrEmptyActorsIDs               = errors.New("empty actors_ids array provided")
)
