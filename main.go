package main

import (
	"github.com/ivan-iver/diaspora/lib"
	"log"
	"os"
	"runtime/debug"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("| Fatal Error | %v", r)
			log.Printf("| Fatal Error | Stack:  %v", string(debug.Stack()))
		}
	}()

	app := lib.NewApp()

	if app.Debug {
		log.Printf(" | Monitor | Ejecutando: %v", app.Version)
	}
	app.Parse(os.Args[1:])
	if app.Debug {
		log.Printf(" | Monitor | Finalize: %v", app.Version)
	}
}
