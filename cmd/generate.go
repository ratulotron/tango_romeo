package cmd

import (
	"log"

	"github.com/ratulotron/tango_romeo/internal/app"
	"github.com/spf13/cobra"
)

func newGenerateCmd(app *app.App) *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "'tango_romeo generate' creates the tech radar file.",
		Long: `'tango_romeo generate' takes the radar file and generates the 
		tech radar file. Eg.: 
		tango_romeo generate -r radar_example.yml -t template/index.gohtml -o output/index.html`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.LoadRadarFromFile(); err != nil {
				log.Println("Can't load tech radar data from file: ", err)
				return
			}

			log.Println("Technologies in radar: ", len(app.Data.Technologies))
			if err := app.GenerateRadar(); err != nil {
				log.Println("Tech radar generate failed: ", err)
				return
			}

		},
	}

	generateCmd.PersistentFlags().StringVarP(
		&app.RadarFile, "radar", "r", "example_radar.yml",
		"Path to tech radar data yml file.",
	)
	generateCmd.PersistentFlags().StringVarP(
		&app.TemplateFile, "template", "t", "template/index.gohtml",
		"Path to tech radar data yml file.",
	)
	generateCmd.PersistentFlags().StringVarP(
		&app.OutputFile, "output", "o", "output/index.html",
		"Path to tech radar data yml file.",
	)

	return generateCmd

}

// var radarFile, templateFile, outputFile string

// // Get command line flag
// var generateCmd = &cobra.Command{
// 	Use:   "tango_romeo",
// 	Short: "`tango_romeo` generates a tech radar.",
// 	Long:  `tango_romeo generates a tech radar for your team, from an yml file.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		app := App{
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

// func Execute() {

// 	rootCmd.PersistentFlags().StringVarP(
// 		&radarFile, "radar", "r", "radar.yml",
// 		"Path to tech radar data yml file.",
// 	)
// 	rootCmd.PersistentFlags().StringVarP(
// 		&templateFile, "template", "t", "template/index.html",
// 		"Path to tech radar data yml file.",
// 	)
// 	rootCmd.PersistentFlags().StringVarP(
// 		&outputFile, "output", "o", "output/index.html",
// 		"Path to tech radar data yml file.",
// 	)
// 	cobra.CheckErr(rootCmd.Execute())

// 	return app

// }
