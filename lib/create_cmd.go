package lib

import (
	"github.com/ivan-iver/diaspora/lib/time"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Create struct contains all about migration script creation
type Create struct {
	*time.Serie
}

// NewCreateCommand initialize new Create command.
func NewCreateCommand() (cmd *Create) {
	cmd = &Create{
		Serie: time.NewSerie(),
	}
	return
}

// Run generates a new migration file
func (cmd *Create) Run(c *kingpin.ParseContext) (err error) {
	var serie = cmd.Formated
	log.Info("Serie: ", serie)
	return
}
