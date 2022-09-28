package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/viper"
)

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
	// pwd, _ := os.Getwd()
	// log.Println(pwd)
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

func NewApp() *App {
	app := &App{}
	return app
}
