package lib

import (
	"github.com/ivan-iver/diaspora/lib/time"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Command define the link between parameters and operative logic
type Command struct {
}

//
func (cmd *Command) RunUp(c *kingpin.ParseContext) (err error) {
	if *debug {
		log.Debug("Algo que poner en debug")
	}
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
