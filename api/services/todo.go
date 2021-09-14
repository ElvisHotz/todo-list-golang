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

func (s TodoService) GetById(id uint) (result models.Todos, err error) {

	db := database.GetDB()
	defer db.Close()

	var todo models.Todos
	if erro := repositories.NewTodoRepository(db).GetById(id, &todo); erro != nil {
		return models.Todos{}, erro
	} else {
		return todo, nil
	}
}

func (s TodoService) New(todo models.Todos) (result models.Todos, err error) {

	db := database.GetDB()
	defer db.Close()

	if erro := repositories.NewTodoRepository(db).New(&todo); erro != nil {
		return models.Todos{}, erro
	} else {
		return todo, nil
	}
}

func (s TodoService) Update(id uint, todo models.Todos) (result models.Todos, err error) {

	db := database.GetDB()
	defer db.Close()

	if erro := repositories.NewTodoRepository(db).Update(id, &todo); erro != nil {
		return models.Todos{}, erro
	} else {
		return todo, nil
	}
}

func (s TodoService) Delete(id uint) (err error) {

	db := database.GetDB()
	defer db.Close()

	if erro := repositories.NewTodoRepository(db).Delete(id); erro != nil {
		return erro
	} else {
		return nil
	}
}
