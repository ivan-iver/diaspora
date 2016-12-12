package lib

import (
	"github.com/ivan-iver/diaspora/lib/time"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Up command execute pending migrations
type Up struct {
	*time.Serie
}

// NewUpCommand initialize new Create command.
func NewUpCommand() (cmd *Up) {
	cmd = &Up{
		Serie: time.NewSerie(),
	}
	return
}

// Run execute the migration command
func (cmd *Up) Run(c *kingpin.ParseContext) (err error) {
	log.IfDebug("Algo que poner en debug")
	log.Info("Dentro de RunUp ", cmd)
	return
}
