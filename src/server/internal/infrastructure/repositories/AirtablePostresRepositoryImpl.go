package repositories

import (
	"gorm.io/gorm"
)

func NewAirtableRepository(db *gorm.DB) *airtableRepository {
	return &airtableRepository{db: db}
}

type airtableRepository struct {
	db *gorm.DB
}

func (a *airtableRepository) CreateNewWebhook(){
}
