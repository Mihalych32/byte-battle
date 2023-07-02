package handler

import (
	"byte-battle_backend/internal/entity"
	"byte-battle_backend/internal/repo"
	"byte-battle_backend/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

var (
	ErrDecodeBody             = errors.New("Could not decode body")
	ErrServerError            = errors.New("Server error")
	ErrCredentialsNotProvided = errors.New("Credentials were not provided")
	ErrUsernameNotProvided    = errors.New("Username was not provided")
	ErrEmailNotProvided       = errors.New("Email was not provided")
	ErrPasswordNotProvided    = errors.New("Password was not provided")
)

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {

	var body entity.RegisterUserBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrDecodeBody.Error()})
		return
	}

	if body.Username == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrUsernameNotProvided.Error()})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrEmailNotProvided.Error()})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrPasswordNotProvided.Error()})
		return
	}

	err := h.uc.RegisterUser(&body)
	if err != nil {
		if err == repo.ErrNonUniqueUsername || err == repo.ErrNonUniqueEmail {
			c.JSON(http.StatusConflict, map[string]string{"message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, map[string]string{"message": ErrServerError.Error()})
		}
		return
	}
	c.Status(http.StatusCreated)
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	var body entity.LoginUserBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrDecodeBody.Error()})
		return
	}

	if body.Username == "" && body.Email == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrCredentialsNotProvided.Error()})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusBadRequest, map[string]string{"message": ErrPasswordNotProvided.Error()})
		return
	}

	err := h.uc.LoginUser(&body)
	if err != nil {
		if err == repo.ErrInvalidCredentials {
			c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, map[string]string{"message": ErrServerError.Error()})
		}
		return
	}
	c.Status(http.StatusOK)
}
