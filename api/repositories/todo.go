package repositories

import (
	"todo-list/models"

	"github.com/jinzhu/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (r todoRepository) GetAll(todos *[]models.Todos) error {
	return r.db.Find(&todos).Error
}
