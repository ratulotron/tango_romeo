package main

import (
	"fmt"

	"github.com/ratulotron/tango_romeo/cmd"
	"github.com/ratulotron/tango_romeo/internal/app"
	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Heyo")
	command := cmd.NewRootCmd(&app.App{})
	cobra.CheckErr(command.Execute())
}
