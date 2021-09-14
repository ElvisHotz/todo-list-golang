package controllers

import (
	"todo-list/models"
	"todo-list/services"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
)

type TodoController struct {
	Service services.TodoService
}

func NewTodoController() *TodoController {
	return &TodoController{Service: *services.NewtodoService()}
}

func (c *TodoController) GetAll(ctx iris.Context) {
	todos, err := c.Service.GetAll()
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err.Error()})
		return
	}
	ctx.StopWithJSON(iris.StatusOK, todos)

}

func (c *TodoController) GetById(ctx iris.Context) {
	id := cast.ToUint(ctx.Params().Get("id"))

	todo, err := c.Service.GetById(id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err.Error()})
		return
	}
	ctx.StopWithJSON(iris.StatusOK, todo)

}

func (c *TodoController) New(ctx iris.Context) {
	var todo models.Todos

	err := ctx.ReadJSON(&todo)
	if err != nil {
		ctx.StopWithStatus(iris.StatusUnprocessableEntity)
		return
	}

	todo, err = c.Service.New(todo)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err.Error()})
		return
	}
	ctx.StopWithJSON(iris.StatusCreated, todo)

}

func (c *TodoController) Update(ctx iris.Context) {
	id := cast.ToUint(ctx.Params().Get("id"))
	var todo models.Todos

	err := ctx.ReadJSON(&todo)
	if err != nil {
		ctx.StopWithStatus(iris.StatusUnprocessableEntity)
		return
	}
	todo.ID = id
	todo, err = c.Service.Update(id, todo)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err.Error()})
		return
	}
	ctx.StopWithJSON(iris.StatusOK, todo)

}

func (c *TodoController) Delete(ctx iris.Context) {
	id := cast.ToUint(ctx.Params().Get("id"))
	err := c.Service.Delete(id)
	if err != nil {
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err.Error()})
		return
	}
	ctx.StopWithJSON(iris.StatusOK, iris.Map{"message": "OK"})

}
