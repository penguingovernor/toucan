package commands

import (
	"fmt"
	"io"
	"os"
)

type commandFunc func(Argument) error

// Argument contains
// all information necessary
// to start the encryption/decryption
// process
type Argument struct {
	keyFile  io.ReadCloser
	ivFile   io.ReadCloser
	dataFile io.ReadCloser
	output   io.WriteCloser
	action   commandFunc
}

func closeAllHelper(tClosers ...io.Closer) {
	for _, closer := range tClosers {
		closer.Close()
	}
}

var commandDispacter = map[string]commandFunc{
	"encrypt": encrypt,
	"decrypt": decrypt,
}

// ParseArgs parses arguments in the format of
// Action Message Key IV [Output, if omitted output = stdout]
func ParseArgs(args []string) (Argument, error) {

	minArgs := 4
	maxArgs := 5

	var err error

	// Arg length check.
	if !(minArgs <= len(args) && len(args) <= maxArgs) {
		return Argument{}, fmt.Errorf("insufficient amount of parameters")
	}

	messageFile := args[1]
	keyFile := args[2]
	IVFile := args[3]

	// Handle actions
	action, valid := commandDispacter[args[0]]
	if !valid {
		return Argument{}, fmt.Errorf("unknown action: %s", args[0])
	}
	returnArgs := Argument{}
	returnArgs.action = action

	// Handle output.
	returnArgs.output = os.Stdout
	if len(args) == maxArgs {
		returnArgs.output, err = os.Create(args[4])
		if err != nil {
			return Argument{}, err
		}
	}

	// Handle Message.
	if returnArgs.dataFile, err = os.Open(messageFile); err != nil {
		return Argument{}, err
	}
	// Handle Key.
	if returnArgs.keyFile, err = os.Open(keyFile); err != nil {
		return Argument{}, err
	}
	// Handle IV.
	if returnArgs.ivFile, err = os.Open(IVFile); err != nil {
		return Argument{}, err
	}

	return returnArgs, nil
}

// CloseAll closes all the values in a.
func (a *Argument) CloseAll() {
	closeAllHelper(a.keyFile, a.ivFile, a.dataFile, a.output)
}

// DoAction causes the argument to trigger the action
// as specified internally
func (a Argument) DoAction() error {
	return a.action(a)
}
