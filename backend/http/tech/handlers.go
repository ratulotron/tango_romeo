package tech

import "github.com/kataras/iris/v12"

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
