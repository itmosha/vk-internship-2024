package entity

import "time"

// Actor entity.
type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}

// Actor with all films' ids where actor is present.
// This struct is used in the API response.
type ActorWithFilms struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
	FilmsIDs  []int  `json:"films_ids"`
}

// Actor create body.
type ActorCreateBody struct {
	Name      string `json:"name"`
	Gender    *bool  `json:"gender"`
	BirthDate string `json:"birth_date"`
}

func ValidateActorCreateBody(body *ActorCreateBody) (err error) {
	if len(body.Name) == 0 || len(body.Name) > 100 {
		return ErrInvalidActorNameLength
	}
	if body.Gender == nil {
		return ErrInvalidActorGender
	}
	_, err = time.Parse("01.02.2006", body.BirthDate)
	if err != nil {
		return ErrInvalidActorBirthDate
	}
	return
}

// Actor update body.
type ActorUpdateBody struct {
	Name      string `json:"name"`
	Gender    *bool  `json:"gender"`
	BirthDate string `json:"birth_date"`
}

func ValidateActorUpdateBody(body *ActorUpdateBody) (err error) {
	if len(body.Name) > 100 {
		return ErrInvalidActorNameLength
	}
	if len(body.BirthDate) > 0 {
		_, err = time.Parse("01.02.2006", body.BirthDate)
		if err != nil {
			return ErrInvalidActorBirthDate
		}
	}
	return
}

// Acter replace body.
type ActorReplaceBody struct {
	Name      string `json:"name"`
	Gender    *bool  `json:"gender"`
	BirthDate string `json:"birth_date"`
}

func ValidateActorReplaceBody(body *ActorReplaceBody) (err error) {
	if len(body.Name) == 0 || len(body.Name) > 100 {
		return ErrInvalidActorNameLength
	}
	if body.Gender == nil {
		return ErrInvalidActorGender
	}
	_, err = time.Parse("01.02.2006", body.BirthDate)
	if err != nil {
		return ErrInvalidActorBirthDate
	}
	return
}
