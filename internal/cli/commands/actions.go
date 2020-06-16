package commands

import toucan "github.com/penguingovernor/toucan/crypto"

func encrypt(arg Argument) error {
	err := toucan.Encrypt(arg.dataFile, arg.keyFile, arg.ivFile, arg.output)
	if err != nil {
		return err
	}
	return nil
}

func decrypt(arg Argument) error {
	err := toucan.Decrypt(arg.dataFile, arg.keyFile, arg.ivFile, arg.output)
	if err != nil {
		return err
	}
	return nil
}
