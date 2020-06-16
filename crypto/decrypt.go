package toucan

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"fmt"
	"io"

	"golang.org/x/crypto/sha3"
)

// Decrypt decrypts the cipher with the provided key.
// The initiation vector should only be used once and should be the same as the encryption IV.
// The key should be the only thing that is kept secret.
func Decrypt(data, key, IV io.Reader, output io.Writer) error {
	// Create HMAC and Encryption Key from IV and key.
	cKey, hKey, err := generateKeys(key, IV)
	if err != nil {
		return err
	}

	// Authenticate the data.
	data, err = authenticate(data, hKey)
	if err != nil {
		return err
	}

	// Create RNG from encryption key.
	rng := rngFromKey(cKey)

	// Then decrypt it, if successful.
	// Loop through the data byte by byte, whilst encrypting.
	dataScanner := bufio.NewScanner(data)
	dataScanner.Split(bufio.ScanBytes)
	outputWriter := bufio.NewWriter(output)
	for dataScanner.Scan() {
		msgByte := dataScanner.Bytes()[0]
		rngByte := byte(int8(rng.Int63()))
		outputWriter.WriteByte(msgByte ^ rngByte)
	}
	if err := dataScanner.Err(); err != nil {
		return err
	}

	return outputWriter.Flush()
}

func authenticate(data io.Reader, hKey []byte) (io.Reader, error) {
	// Save the data
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, data); err != nil {
		return nil, err
	}

	// Create the HMAC
	hWriter := hmac.New(sha3.New512, hKey)

	// Check the size.
	if buf.Len() < hWriter.Size() {
		return nil, fmt.Errorf("file does not reach min. size, want %d got %d", hWriter.Size(), buf.Len())
	}

	// Get the bytes, then from those get the msg and MAC.
	dataBytes := buf.Bytes()
	msgBytes := dataBytes[:len(dataBytes)-hWriter.Size()]
	givenMAC := dataBytes[len(dataBytes)-hWriter.Size():]

	// Hash the msg to compute the MAC we want.
	if _, err := hWriter.Write(msgBytes); err != nil {
		return nil, err
	}
	wantMAC := hWriter.Sum(nil)

	// Compare them.
	if !hmac.Equal(givenMAC, wantMAC) {
		return nil, fmt.Errorf("MAC is invalid")
	}

	// Return the msgbytes as a new reader.
	return bytes.NewReader(msgBytes), nil
}
