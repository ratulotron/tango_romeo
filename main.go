package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

type (
	/* Tech Radar */
	Tech struct {
		Name        string `mapstructure:"name"`
		TechType    string `mapstructure:"type"`
		TechStatus  string `mapstructure:"status"`
		Description string `mapstructure:"description"`
		Website     string `mapstructure:"website"`
	}

	Radar struct {
		Types        []string `mapstructure:"types"`
		Statuses     []string `mapstructure:"statuses"`
		Technologies []Tech   `mapstructure:"technologies"`
	}
	/* App config */
	App struct {
		TemplateFile string `mapstructure:"template"`
		OutputFile   string `mapstructure:"output"`
		RadarFile    string `mapstructure:"radar_file"`
		Data         Radar  `mapstructure:"radar_data"`
	}
)

func (a *App) LoadRadarFromFile() error {
	viper.SetConfigName(a.RadarFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&a.Data); err != nil {
		return err
	}

	return nil
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func (a *App) GenerateRadar() error {
	log.Println("Template file: ", a.TemplateFile)
	log.Println("Output file: ", a.OutputFile)
	_, fileName := filepath.Split(a.TemplateFile)
	t, err := template.New(fileName).Funcs(
		template.FuncMap{
			"ToUpper": strings.ToUpper,
		}).ParseFiles(a.TemplateFile)
	// t, err := template.ParseFiles(a.TemplateFile)

	if err != nil {
		log.Fatalf("Can't load template: %s", err)
		return err
	}
	// "./index.html"
	f, err := create(a.OutputFile)
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	if err := t.ExecuteTemplate(f, fileName, a.Data); err != nil {
		log.Println("Failed to generate file: ", err)
		return err
	}

	return nil
}

/*
	Read cli params
*/

func InitApp() App {
	var radarFile, templateFile, outputFile string
	app := App{}
	// Get command line flag
	var rootCmd = &cobra.Command{
		Use:   "tango_romeo",
		Short: "`tango_romeo` generates a tech radar.",
		Long:  `tango_romeo generates a tech radar for your team, from an yml file.`,
		Run: func(cmd *cobra.Command, args []string) {
			app = App{
				RadarFile:    radarFile,
				TemplateFile: templateFile,
				OutputFile:   outputFile,
			}
			if err := app.LoadRadarFromFile(); err != nil {
				log.Println("Can't load tech radar data from file: ", err)
				return
			}
			log.Println("Technologies in radar: ", len(app.Data.Technologies))
		},
	}

	rootCmd.PersistentFlags().StringVarP(
		&radarFile, "radar", "r", "radar.yml",
		"Path to tech radar data yml file.",
	)
	rootCmd.PersistentFlags().StringVarP(
		&templateFile, "template", "t", "template/index.html",
		"Path to tech radar data yml file.",
	)
	rootCmd.PersistentFlags().StringVarP(
		&outputFile, "output", "o", "output/index.html",
		"Path to tech radar data yml file.",
	)
	cobra.CheckErr(rootCmd.Execute())

	return app
}

func main() {
	fmt.Println("Heyo")
	app := InitApp()
	app.GenerateRadar()
	// generateTemplate(radarData)
}
