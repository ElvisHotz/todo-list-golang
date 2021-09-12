package main

import (
	"fmt"
	"todo-list/config"
	"todo-list/database"
	"todo-list/routers"

	"github.com/kataras/iris/v12"
)

func main() {
	println("Elvis hotz")

	config.LoadConfig()
	db := database.GetDB()
	database.IniTables(db)
	defer db.Close()

	app := newApp()
	routers.InitRouter(app)
	err := app.Run(iris.Addr(fmt.Sprintf(":%d", config.ApiPort)), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}

}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)

	app.AllowMethods(iris.MethodOptions)
	return app
}
