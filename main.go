package main

import (
	"log"
	"os"

	"github.com/penguingovernor/toucan/cli/commands"
)

func main() {
	args, err := commands.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	if err := args.DoAction(); err != nil {
		log.Fatalln(err)
	}
	args.CloseAll()
}
