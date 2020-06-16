package toucan

import (
	"encoding/binary"
	"io"
	"math/rand"

	"golang.org/x/crypto/sha3"
)

// generateKeys reads all the information from the key and IV sources
// and returns two derived keys. Each key contains 32 bytes.
func generateKeys(key, IV io.Reader) ([]byte, []byte, error) {
	// Hash the key and IV.
	hashWriter := sha3.New512()
	if _, err := io.Copy(hashWriter, key); err != nil {
		return nil, nil, err
	}
	if _, err := io.Copy(hashWriter, IV); err != nil {
		return nil, nil, err
	}
	sum := hashWriter.Sum(nil)

	// Return the checksum split in two.
	return sum[len(sum)/2:], sum[:len(sum)/2], nil
}

// rngFromKey returns a RNG from the given key.
func rngFromKey(key []byte) *rand.Rand {
	seed := int64(binary.BigEndian.Uint64(key))
	source := rand.NewSource(seed)
	return rand.New(source)
}
