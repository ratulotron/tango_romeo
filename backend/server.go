package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.WriteString(`Root not found handler.
        This will be applied everywhere except the /api/* requests.`)
	})

	api := app.Party("/api")
	{
		v1 := api.Party("/v1")
		{
			v1.Get("/health", Health)
			// GET: http://localhost:8080/api/v1/techs
			v1.Get("/techs", List)
		}
	}

	app.Listen(":8080")
}

func Health(ctx iris.Context) {
	ctx.JSON(&map[string]string{"message": "pong"})
}

// Tech example.
type Tech struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func List(ctx iris.Context) {
	techs := []Tech{
		{"Python", "Python is an interpreted language."},
	}

	ctx.JSON(techs)
}
