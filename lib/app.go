// Package lib contains all application libraries.
package lib

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	appName   = "Diaspora"
	appNum    = "0.1"
	desc      = "A command-line database migration tool"
	createMsg = "Create new migration file with the scaffolding attributes"
	upMsg     = "Migrate the DB to the most recent version available"
	debugMsg  = "Debug mode enabled"
	envMsg    = "Set the environment configuration [dev, production, etc] dev is the default"
)

var hash = ""
var log *Logger

// App is a principal structure where we join all application components.
type App struct {
	app         *kingpin.Application
	Log         *Logger
	Config      *Config
	Environment string
	Version     string
	Name        string
	Debug       bool
	*Command
}

func config() (config *Config) {
	var err error
	if config, err = NewConfig(); err != nil {
		fmt.Errorf("| Error | Reading config file: %v \n", err)
	}
	return
}

// NewApp initialize all App fields
func NewApp() (application *App) {
	var config = config()
	log = GetLogger()
	application = &App{
		Name:    appName,
		Config:  config,
		Log:     log,
		Command: &Command{},
		Version: fmt.Sprintf("%v %v %v : %v", appName, appNum, hash, desc),
	}
	application.init()
	application.app.Version(application.Version)
	application.SetActions()
	return
}

func (a *App) init() {
	a.app = kingpin.New(a.Name, desc)
	a.app.Flag("debug", debugMsg).Short('d').Default("false").BoolVar(&a.Debug)
	a.app.Flag("env", envMsg).Short('e').Default("dev").StringVar(&a.Environment)
}

// SetActions create the relationship between commands, functions and parameters
func (a *App) SetActions() {
	a.app.Command("up", upMsg).Action(a.RunUp)
	a.app.Command("create", createMsg).Action(a.RunCreate)
	//	one.Arg("user_id", "User identifier").Required().Int64Var(&a.UserId)
}

// Parse all command line arguments
func (a *App) Parse(args []string) {
	kingpin.MustParse(a.app.Parse(args))
}
