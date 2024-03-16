package entity

import "time"

// Actor entity.
type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}

// Actor create body.
type ActorCreateBody struct {
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}

func ValidateActorCreateBody(body *ActorCreateBody) (err error) {
	if len(body.Name) == 0 || len(body.Name) > 100 {
		return ErrInvalidActorNameLength
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
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}

// Acter replace body.
type ActorReplaceBody struct {
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}
