package entities

import (
	"fmt"

	"github.com/emperorsixpacks/dailbot/src/pkg/utils"
	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
	secret    string
}

func (u User) FullName() string {
	return fmt.Sprintf("%v %v", u.firstName, u.lastName)
}

func (u User) Email() string {
	return u.email
}
func (u *User) GenerateUserSecrete() string {
	str, err := utils.GenerateRandomString(16)
	if err != nil {
		panic(err)
	}
	u.secret = str
	return u.secret
}

