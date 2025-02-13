package repositories

import (
	"github.com/emperorsixpacks/dailbot/src/internal/domain/entities"
	"gorm.io/gorm"
)

func NewPostgresRepository[T any](db *gorm.DB) *postgresRepository[T] {
	return &postgresRepository[T]{db: db}
}

type postgresRepository[T any] struct {
	db *gorm.DB
}

func (p postgresRepository[T]) Create(enty *entities.AirtableWebhook) T {
	var entity T
	model := enty.Model()
	p.db.Create(model)
	return entity
}
