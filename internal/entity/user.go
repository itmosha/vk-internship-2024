package entity

// User entity.
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

// Simplified user register body, no password.
// The role can be changed in the database.
type UserRegisterBody struct {
	Username string `json:"username"`
}

func ValidateUserRegisterBody(body *UserRegisterBody) (err error) {
	if len(body.Username) == 0 || len(body.Username) > 100 {
		err = ErrInvalidUsernameLength
	}
	return
}

// Simplified user login body, no password.
type UserLoginBody struct {
	Username string `json:"username"`
}

func ValidateUserLoginBody(body *UserLoginBody) (err error) {
	if len(body.Username) == 0 || len(body.Username) > 100 {
		err = ErrInvalidUsernameLength
	}
	return
}
