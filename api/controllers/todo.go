package controllers

import (
	"todo-list/models"
	"todo-list/services"

	"github.com/kataras/iris/v12"
)

type TodoController struct {
	Service services.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{Service: *services.NewtodoService()}
}

func (c *TodoController) GetAllMVC() (result []models.Todos, err error) {
	return c.Service.GetAll()

}

func (c *TodoController) GetAll(ctx iris.Context) {
	todos, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(err)
	}
	ctx.JSON(todos)

}
