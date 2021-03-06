package main

import (
	"fmt"
	"os"

	"github.com/ConradIrwin/font/commands"
)

func main() {

	command := "help"

	if len(os.Args) > 1 {
		command = os.Args[1]
		os.Args = os.Args[1:]
	}

	switch command {
	case "scrub":
		commands.Scrub()
	case "info":
		commands.Info()
	case "stats":
		commands.Stats()
	case "metrics":
		commands.Metrics()
	default:
		fmt.Println(`
Usage: font [scrub|info|stats] font.[otf,ttf,woff]

info: prints the name table (contains metadata)
metrics: prints the hhea table (contains font metrics)
stats: prints each table and the amount of space used
scrub: remove the name table (saves significant space)`)
	}

}
