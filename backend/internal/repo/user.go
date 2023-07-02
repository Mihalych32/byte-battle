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
	ErrServerError        = errors.New("Server error")
	ErrNonUniqueUsername  = errors.New("Non unique username")
	ErrNonUniqueEmail     = errors.New("Non unique email")
	ErrInvalidCredentials = errors.New("Invalid credentials")
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

func (r *UserRepo) CheckUserCredentials(body *entity.CheckUserCredentialsBody) (user *entity.User, err error) {

	var credsOk int
	if body.Email == "" {
		query := `
			SELECT COUNT(*)
			FROM users
			WHERE username=$1
			AND encrypted_pwd=$2;`
		err = r.store.DB.
			QueryRow(query, body.Username, body.EncryptedPwd).
			Scan(&credsOk)
	} else {
		query := `
			SELECT COUNT(*)
			FROM users
			WHERE email=$1
			AND encrypted_pwd=$2;`
		err = r.store.DB.
			QueryRow(query, body.Email, body.EncryptedPwd).
			Scan(&credsOk)
	}

	if err != nil {
		err = ErrServerError
		return
	}

	if credsOk == 1 {
		user = &entity.User{}

		if body.Email == "" {
			userQuery := `
				SELECT id, username, email, role, encrypted_pwd
				FROM users
				WHERE username=$1;`
			err = r.store.DB.
				QueryRow(userQuery, body.Username).
				Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)
		} else {
			userQuery := `
				SELECT id, username, email, role, encrypted_pwd
				FROM users
				WHERE email=$1;`
			err = r.store.DB.
				QueryRow(userQuery, body.Email).
				Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.EncryptedPwd)
		}

		if err != nil {
			err = ErrServerError
			return
		}
	} else {
		err = ErrInvalidCredentials
	}
	return
}
