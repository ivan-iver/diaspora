package lib

import (
	"diaspora/lib/time"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Command define the link between parameters and operative logic
type Command struct {
}

//
func (cmd *Command) RunUp(c *kingpin.ParseContext) (err error) {
	log.Info("Dentro de RunUp ", cmd)
	return
}

func (cmd *Command) RunCreate(c *kingpin.ParseContext) (err error) {
	var serie *time.Serie
	if serie, err = time.NewSerie(); err != nil {
		log.Error("On NewSerie")
		return
	}
	log.Info("Serie: ", serie.Formated)
	return
}
