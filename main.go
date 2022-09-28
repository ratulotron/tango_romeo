package main

import (
	"fmt"

	"github.com/ratulotron/tango_romeo/cmd"
	"github.com/ratulotron/tango_romeo/internal/app"
	"github.com/spf13/cobra"
)

/*
	Types:
		- Datastore
		- Data Management
		- Infrastructure
		- Languages
	Statuses:
		- adopt
		- assess
		- trial
		- hold
*/

/*
	Read cli params
*/

// func InitApp() App {
// var radarFile, templateFile, outputFile string
// app := App{}
// // Get command line flag
// var rootCmd = &cobra.Command{
// 	Use:   "tango_romeo",
// 	Short: "`tango_romeo` generates a tech radar.",
// 	Long:  `tango_romeo generates a tech radar for your team, from an yml file.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		app = App{
// 			RadarFile:    radarFile,
// 			TemplateFile: templateFile,
// 			OutputFile:   outputFile,
// 		}
// 		if err := app.LoadRadarFromFile(); err != nil {
// 			log.Println("Can't load tech radar data from file: ", err)
// 			return
// 		}
// 		log.Println("Technologies in radar: ", len(app.Data.Technologies))
// 	},
// }

// rootCmd.PersistentFlags().StringVarP(
// 	&radarFile, "radar", "r", "radar.yml",
// 	"Path to tech radar data yml file.",
// )
// rootCmd.PersistentFlags().StringVarP(
// 	&templateFile, "template", "t", "template/index.html",
// 	"Path to tech radar data yml file.",
// )
// rootCmd.PersistentFlags().StringVarP(
// 	&outputFile, "output", "o", "output/index.html",
// 	"Path to tech radar data yml file.",
// )
// cobra.CheckErr(rootCmd.Execute())

// return app
// }

func main() {
	fmt.Println("Heyo")
	command := cmd.NewRootCmd(&app.App{})
	cobra.CheckErr(command.Execute())
	// app := InitApp()
	// app.GenerateRadar()
	// generateTemplate(radarData)
}
