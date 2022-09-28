package cmd

import (
	"github.com/ratulotron/tango_romeo/internal/app"
	"github.com/spf13/cobra"
)

func NewRootCmd(app *app.App) *cobra.Command {
	// Get command line flag
	rootCmd := &cobra.Command{
		Use:   "tango_romeo",
		Short: "`tango_romeo` generates a tech radar.",
		Long:  `tango_romeo generates a tech radar for your team, from an yml file.`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	rootCmd.AddCommand(newGenerateCmd(app))
	return rootCmd
}
