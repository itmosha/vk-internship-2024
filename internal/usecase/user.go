package usecase

import (
	"context"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	jwtfuncs "github.com/itmosha/vk-internship-2024/pkg/jwt_funcs"
)

type UserRepoInterface interface {
	Insert(ctx *context.Context, receivedUser *entity.User) (createdUser *entity.User, err error)
	SelectByUsername(ctx *context.Context, username string) (user *entity.User, err error)
}

type UserUsecase struct {
	userRepo UserRepoInterface
}

func NewUserUsecase(userRepo UserRepoInterface) *UserUsecase {
	return &UserUsecase{userRepo}
}

func (u *UserUsecase) Register(ctx *context.Context, body *entity.UserRegisterBody) (createdUser *entity.User, err error) {
	userToCreate := &entity.User{
		Username: body.Username,
	}
	createdUser, err = u.userRepo.Insert(ctx, userToCreate)
	return
}

func (u *UserUsecase) Login(ctx *context.Context, body *entity.UserLoginBody) (accessToken string, err error) {
	user, err := u.userRepo.SelectByUsername(ctx, body.Username)
	if err != nil {
		return
	}
	accessToken, err = jwtfuncs.CreateAccessToken(&jwtfuncs.AccessTokenClaims{
		ID:       user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	})
	return
}
