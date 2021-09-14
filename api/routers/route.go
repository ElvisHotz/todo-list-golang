package routers

import (
	"todo-list/controllers"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {

	bathUrl := "/api"

	todos := app.Party(bathUrl + "/todos")
	todos.Get("/", controllers.NewTodoController().GetAll)
	todos.Get("/{id}", controllers.NewTodoController().GetById)
	todos.Post("/", controllers.NewTodoController().New)
	todos.Put("/{id}", controllers.NewTodoController().Update)
	todos.Delete("/{id}", controllers.NewTodoController().Delete)

}
