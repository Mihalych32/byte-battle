package usecase

import (
	"byte-battle_backend/internal/entity"
	"byte-battle_backend/internal/repo"
	"crypto/sha256"
	"encoding/hex"
)

type UserUsecase struct {
	repo *repo.UserRepo
}

func NewUserUsecase(repo *repo.UserRepo) *UserUsecase {
	return &UserUsecase{repo}
}

func (u *UserUsecase) RegisterUser(body *entity.RegisterUserBody) (err error) {
	h := sha256.New()
	h.Write([]byte(body.Password))
	encryptedPwd := hex.EncodeToString(h.Sum(nil))

	createUserBody := &entity.CreateUserBody{
		Username:     body.Username,
		Email:        body.Email,
		EncryptedPwd: encryptedPwd,
	}
	_, err = u.repo.CreateUser(createUserBody)
	return
}
