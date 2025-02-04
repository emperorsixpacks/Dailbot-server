package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type baseModel struct {
	gorm.Model
}

type UserModel struct {
	baseModel
	ID        uuid.UUID `gorm:"primaryKey;column:id"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
}

type WebhookModel struct {
	baseModel
	ID            uuid.UUID `gorm:"column:id"`
	WebHookID     string    `gorm:"column:webhook_id"`
	BaseID        string    `gorm:"column:base_id;uniqueIndex"`
	PersonalToken string    `gorm:"column:personal_token;uniqueIndex"`
	Cursor        string    `gorm:"column:cursor"`
	TablesIDs     []string  `gorm:"column:tables;type:string[]"`
	WebHookSecret string    `gorm:"column:webhook_secret;uniqueIndex"`
}

func (UserModel) TableName() string {
	return "airtable_webhook"
}

func (WebhookModel) TableName() string {
	return "users"
}
