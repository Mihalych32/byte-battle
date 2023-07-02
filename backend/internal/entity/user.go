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
