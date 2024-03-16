package entity

import "time"

// Film entity.
type Film struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      int    `json:"rating"`
}

// Film create body.
type FilmCreateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      int    `json:"rating"`
	ActorsIDs   []int  `json:"actors_ids"`
}

func ValidateFilmCreateBody(body *FilmCreateBody) (err error) {
	if len(body.Title) == 0 || len(body.Title) > 150 {
		return ErrInvalidFilmTitleLength
	}
	if len(body.Description) > 1000 {
		return ErrInvalidFilmDescriptionLength
	}
	_, err = time.Parse("01.02.2006", body.ReleaseDate)
	if err != nil {
		return ErrInvalidFilmReleaseDate
	}
	if body.Rating < 0 || body.Rating > 10 {
		return ErrInvalidFilmRating
	}
	if len(body.ActorsIDs) == 0 {
		return ErrEmptyActorsIDs
	}
	return
}

// Film update body.
type FilmUpdateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      *int   `json:"rating"`
}

func ValidateFilmUpdateBody(body *FilmUpdateBody) (err error) {
	if len(body.Title) > 150 {
		return ErrInvalidFilmTitleLength
	}
	if len(body.Description) > 1000 {
		return ErrInvalidFilmDescriptionLength
	}
	if len(body.ReleaseDate) > 0 {
		_, err = time.Parse("01.02.2006", body.ReleaseDate)
		if err != nil {
			return ErrInvalidFilmReleaseDate
		}
	}
	if body.Rating != nil && (*body.Rating < 0 || *body.Rating > 10) {
		return ErrInvalidFilmRating
	}
	return
}

// Film replace body.
type FilmReplaceBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      int    `json:"rating"`
}

// Film sort params.
type FilmSortParams struct {
	Field string
	Order string
}

// Fields that Films can be sorted by.
var FilmSortFields = [...]string{"title", "rating", "release_date"}

// Directions that Films can be sorted in.
var FilmSortOrder = [...]string{"asc", "desc"}

// Fields that Films can be searched by.
var FilmSearchFields = [...]string{"title", "actor_name"}

// Default sort field.
const FilmDefaultSortField = "rating"

// Default sort order.
const FilmDefaultSortOrder = "desc"

// IsValidParam checks if param is valid.
func IsValidParam(param, value string) bool {
	switch param {
	case "sort_field":
		for _, f := range FilmSortFields {
			if f == value {
				return true
			}
		}
	case "sort_order":
		for _, d := range FilmSortOrder {
			if d == value {
				return true
			}
		}
	case "search_field":
		for _, f := range FilmSearchFields {
			if f == value {
				return true
			}
		}
	}
	return false
}
