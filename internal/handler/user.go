package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/entity"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
	"github.com/itmosha/vk-internship-2024/pkg/logger"
)

type UserUsecaseInterface interface {
	Register(ctx *context.Context, body *entity.UserRegisterBody) (user *entity.User, err error)
	Login(ctx *context.Context, body *entity.UserLoginBody) (accessToken string, err error)
}

type UserHandler struct {
	userUsecase UserUsecaseInterface
	logger      *logger.Logger
}

func NewUserHandler(userUsecase UserUsecaseInterface, logger *logger.Logger) *UserHandler {
	return &UserHandler{userUsecase, logger}
}

func (h *UserHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			h.logger.Log(r, http.StatusBadRequest, ErrEmptyBody)
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.UserRegisterBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateUserRegisterBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		film, err := h.userUsecase.Register(&ctx, body)
		if err != nil {
			switch err {
			case repo.ErrNonUniqueUsername:
				h.logger.Log(r, http.StatusConflict, err)
				returnError(w, http.StatusConflict, err)
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(film)
		h.logger.Log(r, http.StatusCreated, nil)
	}
}

func (h *UserHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isEmptyBody(r) {
			h.logger.Log(r, http.StatusBadRequest, ErrEmptyBody)
			returnError(w, http.StatusBadRequest, ErrEmptyBody)
			return
		}
		body, err := readBodyToStruct(r, &entity.UserLoginBody{})
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}
		err = entity.ValidateUserLoginBody(body)
		if err != nil {
			h.logger.Log(r, http.StatusBadRequest, err)
			returnError(w, http.StatusBadRequest, err)
			return
		}

		ctx := context.Background()
		accessToken, err := h.userUsecase.Login(&ctx, body)
		if err != nil {
			switch err {
			case repo.ErrUserNotFound:
				h.logger.Log(r, http.StatusBadRequest, err)
				returnError(w, http.StatusBadRequest, err)
			default:
				h.logger.Log(r, http.StatusInternalServerError, err)
				returnError(w, http.StatusInternalServerError, ErrServerError)
			}
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
		h.logger.Log(r, http.StatusOK, nil)
	}
}
