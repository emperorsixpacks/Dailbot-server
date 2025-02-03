package entities

import (
	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
}

type BaseUser struct {
	User
	password string
}

type UserRepository interface {
	Createuser(User, string) User
	GetUser(uuid.UUID) User
	VerifyPassword(string) bool
}
