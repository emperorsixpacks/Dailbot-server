package models

import "github.com/google/uuid"

type UserModel struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	password  string
}
