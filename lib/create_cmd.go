package lib

import (
	"github.com/ivan-iver/diaspora/lib/io"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Create struct contains all about migration script creation
type Create struct {
	Name string
}

// Run generates a new migration file
func (cmd *Create) Run(c *kingpin.ParseContext) (err error) {
	tmp, err := io.NewTemplate()
	if err != nil {
		log.Error(err)
		return
	}
	tmp.Name = cmd.Name
	tmp.Path = *dbpath
	err = tmp.Save()
	return
}
