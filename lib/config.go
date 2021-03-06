package lib

import (
	"os"
	"path"
	"strings"

	conf "bitbucket.org/ivan-iver/config"
)

var (
	production = `production`
	filename   = `db.conf`
)

// Config models an entity config reader
type Config struct {
	Pwd          string
	Filename     string
	IsProduction bool
	*conf.Config
}

// NewConfig creates a config struct.
func NewConfig() (config *Config) {
	var err error
	config = &Config{Filename: filename}
	if config.Pwd, err = os.Getwd(); err != nil {
		log.Error(err)
		//		panic(err)
		config.setDefault()
		return
	}

	var file = config.File()
	//	log.Debug("App | Config will be loaded from %v \n", file)
	if config.Config, err = conf.ReadDefault(file); err != nil {
		log.Debug(err)
		config.setDefault()
		return
	}
	//	log.Debug("App | Config loaded successfully! \n")
	config.IsProduction = strings.EqualFold(config.Default("env"), production)
	//log.Info("App | Config.IsProduction ", config.IsProduction)
	return
}

// Set default values when config file does not exists
func (c *Config) setDefault() {
	log.Debug("Using default values ...")
	c.Config = &conf.Config{}
	c.IsProduction = false
}

// File function  returns configuration file path
func (c *Config) File() (file string) {
	return path.Join(c.Pwd, c.Filename)
}

// Default function gets config property from default section
func (c *Config) Default(property string) (result string) {
	var err error
	//log.Printf("App | Property: %v \n", property)
	if result, err = c.String("default", property); err != nil {
		log.Fatalf("| Error | %v \n", err)
		return ""
	}
	//log.Printf("App | Value: %v \n", result)
	return
}

// StringDefault gets config property from default section or use default value
func (c *Config) StringDefault(property string, strDefault string) (result string) {
	var err error
	//log.Printf("Config | Property: %v", property)
	if result, err = c.String("default", property); err != nil {
		//		log.Fatalf("| Error | %v \n", err)
		return strDefault
	}
	//log.Printf("Config | Value: %v \n", result)
	return
}

// BooleanDefault function gets config property from default section or use default boolean value
func (c *Config) BooleanDefault(property string, boolDefault bool) (result bool) {
	var err error
	// log.Printf("Config | Property: %v", property)
	if result, err = c.Bool("default", property); err != nil {
		log.Fatalf("| Error | %v \n", err)
		return boolDefault
	}
	//log.Printf("Config | Value: %v \n", result)
	return
}
