package services

import (
	"todo-list/database"
	"todo-list/models"
	"todo-list/repositories"
)

type TodoService struct {
}

func NewtodoService() *TodoService {
	return &TodoService{}
}

func (s TodoService) GetAll() (result []models.Todos, err error) {

	db := database.GetDB()
	defer db.Close()

	var todos []models.Todos
	if erro := repositories.NewTodoRepository(db).GetAll(&todos); erro != nil {
		return nil, erro
	} else {
		return todos, nil
	}

}
