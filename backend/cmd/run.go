package cmd

import (
	tech_handlers "tango_romeo/http/tech"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

var App *iris.Application

func init() {
	App = iris.New()

	App.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.WriteString(`Something went wrong`)
	})

	api := App.Party("/api")
	{
		v1 := api.Party("/v1")
		{
			v1.Get("/health", tech_handlers.Health)
			// GET: http://localhost:8080/api/v1/techs
			v1.Get("/techs", tech_handlers.List)
		}
	}
}

// rootCmd represents the base command when called without any subcommands
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run application server",
	Run: func(cmd *cobra.Command, args []string) {
		App.Listen(":8080")
	},
}
