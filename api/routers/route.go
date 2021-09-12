package routers

import (
	"todo-list/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {

	bathUrl := "/api"
	mvc.New(app.Party(bathUrl + "/todos2")).Handle(controllers.NewTodoController())
	app.Handle("GET", bathUrl+"/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	todos := app.Party(bathUrl + "/todos")
	todos.Get("/", controllers.NewTodoController().GetAll)

}
