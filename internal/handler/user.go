package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
)

type UserUsecaseInterface interface {
	Register(ctx *context.Context, body *entity.UserRegisterBody) (user *entity.User, err error)
	Login(ctx *context.Context, body *entity.UserLoginBody) (accessToken string, err error)
}

type UserHandler struct {
	userUsecase UserUsecaseInterface
}

func NewUserHandler(userUsecase UserUsecaseInterface) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.UserRegisterBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateUserRegisterBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		film, err := h.userUsecase.Register(&ctx, body)
		if err != nil {
			switch err {
			case repo.ErrUserNotFound:
				returnError(w, http.StatusBadRequest, err)
				return
			default:
				returnError(w, http.StatusInternalServerError, ErrServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(film)
	}
}

func (h *UserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.UserLoginBody{})
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateUserLoginBody(body)
		if err != nil {
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		accessToken, err := h.userUsecase.Login(&ctx, body)
		if err != nil {
			switch err {
			case repo.ErrUserNotFound:
				returnError(w, http.StatusBadRequest, err)
				return
			default:
				returnError(w, http.StatusInternalServerError, ErrServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
	}
}
