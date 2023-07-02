package entity

type Role int8

const (
	quest Role = iota
	user
	moderator
	admin
)

type User struct {
	ID           int
	Username     string
	Email        string
	Role         int8
	EncryptedPwd string
}

type RegisterUserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserBody struct {
	Username     string
	Email        string
	EncryptedPwd string
}

type CheckUserCredentialsBody struct {
	Username     string
	Email        string
	EncryptedPwd string
}
