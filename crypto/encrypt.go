package toucan

import (
	"bufio"
	"crypto/hmac"
	"fmt"
	"io"
	"math/rand"

	"golang.org/x/crypto/sha3"
)

// Encrypt encrypts the message with the provided key.
// The initiation vector should only be used once.
// The key should be the only thing that is kept secret.
func Encrypt(data, key, IV io.Reader, output io.Writer) error {

	// Create HMAC and Encryption Key from IV and key.
	cKey, hKey, err := generateKeys(key, IV)
	if err != nil {
		return err
	}

	// Create RNG from encryption key.
	rng := rngFromKey(cKey)

	// Then encrypt it.
	return encryptThenMAC(data, output, rng, hKey)
}

// encryptThenMAC encrypts the given data using the given RNG and HMAC key (hKey).
// It then appends the HMAC of that encrypted data to the cipher text.
func encryptThenMAC(data io.Reader, dest io.Writer, rng *rand.Rand, hKey []byte) error {

	// Make two writers:
	// One for the HMAC
	// Another for the HMAC & dest.
	hWriter := hmac.New(sha3.New512, hKey)
	mWriter := bufio.NewWriter(io.MultiWriter(dest, hWriter))

	// Loop through the data byte by byte, whilst encrypting.
	dataScanner := bufio.NewScanner(data)
	dataScanner.Split(bufio.ScanBytes)
	for dataScanner.Scan() {
		msgByte := dataScanner.Bytes()[0]
		rngByte := byte(int8(rng.Int63()))
		mWriter.WriteByte(msgByte ^ rngByte)
	}
	if err := dataScanner.Err(); err != nil {
		return err
	}
	if err := mWriter.Flush(); err != nil {
		return err
	}

	// Add the HMAC bytes.
	hmacBytes := hWriter.Sum(nil)
	bytesWritten, err := dest.Write(hmacBytes)
	if err != nil {
		return err
	}
	if bytesWritten != hWriter.Size() {
		return fmt.Errorf("failed to write all HMAC bytes, aborting")
	}

	return nil
}
