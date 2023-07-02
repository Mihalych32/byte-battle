package repo

import (
	"byte-battle_backend/internal/entity"
	"byte-battle_backend/pkg/postgres"
	"errors"

	"github.com/lib/pq"
)

type UserRepo struct {
	store *postgres.Postgres
}

var (
	ErrServerError       = errors.New("Server error")
	ErrNonUniqueUsername = errors.New("Non unique username")
	ErrNonUniqueEmail    = errors.New("Non unique email")
)

func NewUserRepo(p *postgres.Postgres) *UserRepo {
	return &UserRepo{p}
}

func (r *UserRepo) CreateUser(body *entity.CreateUserBody) (user *entity.User, err error) {
	query := `
		INSERT INTO users (username, email, role, encrypted_pwd)
		VALUES ($1, $2, 1, $3)
		RETURNING id, username, email, role, encrypted_pwd;`

	user = &entity.User{}

	err = r.store.DB.
		QueryRow(query, body.Username, body.Email, body.EncryptedPwd).
		Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				if pqErr.Constraint == "users_username_key" {
					err = ErrNonUniqueUsername
					return
				} else if pqErr.Constraint == "users_email_key" {
					err = ErrNonUniqueEmail
					return
				}
			}
		}
		err = ErrServerError
		return
	}
	return
}
