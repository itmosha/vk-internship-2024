package entity

// Actor entity.
type Actor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}

// Actor create body.
type ActorCreateBody struct {
	Name      string `json:"name" binding:"required"`
	Gender    bool   `json:"gender" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}

// Actor update body.
type ActorUpdateBody struct {
	Name      string `json:"name"`
	Gender    bool   `json:"gender"`
	BirthDate string `json:"birth_date"`
}

// Acter replace body.
type ActorReplaceBody struct {
	Name      string `json:"name" binding:"required"`
	Gender    bool   `json:"gender" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}
