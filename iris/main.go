package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main()  {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())


	app.Handle("GET", "/", func(context iris.Context) {
		context.HTML("<h1>Welcome GoLang World</h1>")
	})

	app.Get("/ping", func(context iris.Context) {
		context.WriteString("pong")
	})

	app.Get("/hello", func(context iris.Context) {
		context.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
