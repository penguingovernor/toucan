package toucan

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"hash/fnv"
	"math/rand"
)

// Encrypt encrypts the message with the provided key.
// The initiation vector should only be used once.
// The key should be the only thing that is kept secret.
func Encrypt(message, key, IV []byte) ([]byte, error) {
	// Create a seed using the IV and key.
	hashFunc := fnv.New64()
	if _, err := hashFunc.Write(append(IV, key...)); err != nil {
		return nil, err
	}
	var seed int64 = int64(hashFunc.Sum64())

	// Use this seed to make a rng.
	rng := rand.NewSource(seed)

	// Create the cipher text.
	cipher := make([]byte, len(message))
	for i, msgByte := range message {
		var randNum byte = byte(int8(rng.Int63()))
		cipher[i] = msgByte ^ randNum
	}

	// Make an HMAC using the key.
	hmacFunc := hmac.New(sha256.New, key)
	if _, err := hmacFunc.Write(cipher); err != nil {
		return nil, err
	}

	// Append the hmac to the cipher
	return append(hmacFunc.Sum(nil), cipher...), nil
}

// Decrypt decrypts the cipher with the provided key.
// The initiation vector should only be used once and should be the same as the encryption IV.
// The key should be the only thing that is kept secret.
func Decrypt(cipher, key, IV []byte) ([]byte, error) {
	// Get the HMAC and remove it from the cipher.
	hmacFunc := hmac.New(sha256.New, key)
	want := cipher[:hmacFunc.Size()]
	cipher = cipher[hmacFunc.Size():]

	// Generate our own HMAC.
	if _, err := hmacFunc.Write(cipher); err != nil {
		return nil, err
	}
	got := hmacFunc.Sum(nil)

	// Validate the HMAC.
	if !hmac.Equal(want, got) {
		return nil, errors.New("decryption failed")
	}

	// Create a seed using the IV and key.
	hashFunc := fnv.New64()
	if _, err := hashFunc.Write(append(IV, key...)); err != nil {
		return nil, err
	}
	var seed int64 = int64(hashFunc.Sum64())

	// Use this seed to make a rng.
	rng := rand.NewSource(seed)

	// Create the plain text.
	message := make([]byte, len(cipher))
	for i, msgByte := range cipher {
		var randNum byte = byte(int8(rng.Int63()))
		message[i] = msgByte ^ randNum
	}

	return message, nil
}
