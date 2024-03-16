package entity

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
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Rating      int    `json:"rating"`
}

// Film update body.
type FilmUpdateBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      int    `json:"rating"`
}

// Film replace body.
type FilmReplaceBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date" binding:"required"`
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
