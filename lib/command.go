package lib

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

type Command struct {
	Log *Logger
}

func (cmd *Command) RunUp(c *kingpin.ParseContext) (err error) {
	log.Printf("Dentro de RunUp %#v", cmd)
	return
}
