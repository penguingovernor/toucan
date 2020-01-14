package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/penguingovernor/toucan/cli/commands"
)

const usage = `toucan - a simple symmetric stream cipher.

FOR EDUCATIONAL PURPOSES ONLY.

Usage:
	toucan [command] msgFile keyFile IVFile [outputFile]

Available Commands:
	encrypt		Encrypt files
	decrypt		Decrypt files

Flags:
	-h, --help		print this help message.

Notes:
	If [outputFile is committed] then stdout is used.
`

func main() {

	// Define the usage func.
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	// If there's no real args then
	// print the usage and exit.
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	// Handle the --help, -h flags.
	flag.Parse()

	// Start the encryption/decryption process.
	args, err := commands.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	if err := args.DoAction(); err != nil {
		log.Fatalln(err)
	}
	args.CloseAll()
}
