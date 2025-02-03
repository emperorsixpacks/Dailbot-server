package entities

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
}

func (u User) FullName() string {
	return fmt.Sprintf("%v %v", u.firstName, u.lastName)
}

func (u User) ToJson() interface{}

func (u User) Email() string {
	return u.email
}

type UserRepository interface {
	CreateUser(User, string) User
	GetUserByID(uuid.UUID) User
	VerifyPassword(string) bool
}
