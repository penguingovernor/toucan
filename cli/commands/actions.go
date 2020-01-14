package commands

import (
	"io/ioutil"

	"github.com/penguingovernor/toucan/crypto/toucan"
)

type cryptoPacket struct {
	IV   []byte
	Key  []byte
	Data []byte
}

func getData(arg Argument) (cryptoPacket, error) {
	dataBytes, err := ioutil.ReadAll(arg.dataFile)
	if err != nil {
		return cryptoPacket{}, err
	}

	IVBytes, err := ioutil.ReadAll(arg.ivFile)
	if err != nil {
		return cryptoPacket{}, err
	}

	KeyBytes, err := ioutil.ReadAll(arg.keyFile)
	if err != nil {
		return cryptoPacket{}, err
	}

	return cryptoPacket{
		IV:   IVBytes,
		Key:  KeyBytes,
		Data: dataBytes,
	}, nil
}

func encrypt(arg Argument) error {
	cp, err := getData(arg)
	if err != nil {
		return err
	}
	cipher, err := toucan.Encrypt(cp.Data, cp.Key, cp.IV)
	if err != nil {
		return err
	}
	if _, err := arg.output.Write(cipher); err != nil {
		return err
	}
	return nil
}

func decrypt(arg Argument) error {
	cp, err := getData(arg)
	if err != nil {
		return err
	}
	pt, err := toucan.Decrypt(cp.Data, cp.Key, cp.IV)
	if err != nil {
		return err
	}
	if _, err := arg.output.Write(pt); err != nil {
		return err
	}
	return nil
}
