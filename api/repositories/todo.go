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

func (r todoRepository) GetById(id uint, todo *models.Todos) error {
	return r.db.Where("id = ?", id).First(&todo).Error
}

func (r todoRepository) New(todo *models.Todos) error {
	return r.db.Create(&todo).Error

}

func (r todoRepository) Update(id uint, todo *models.Todos) error {
	return r.db.Model(&todo).Where("id = ?", id).Update(&todo).Error
}

func (r todoRepository) Delete(id uint) error {
	var todo models.Todos
	err := r.GetById(id, &todo)
	if err != nil {
		return err
	}
	return r.db.Where("id = ?", id).Delete(&todo).Error
}
