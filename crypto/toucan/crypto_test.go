package toucan

import (
	"bytes"
	"encoding/base64"
	"testing"
)

func TestEncrypt(t *testing.T) {
	msg := []byte("Hello World\n")

	t.Run("successful encryption", func(t *testing.T) {
		ct, err := Encrypt(msg, msg, msg)
		if err != nil {
			t.Fatalf("wanted err == nil, got err = %v", err)
		}
		want := "efrjHxns8qKwYtQJolQZdv6+HNRb1teO8xcxM1/3/LKJhsjvfYsXPOZLjXY="
		wantCT, err := base64.StdEncoding.DecodeString(want)
		if err != nil {
			t.Fatalf("wanted err == nil, got err = %v", err)
		}
		if !bytes.Equal(wantCT, ct) {
			t.Fatalf("want != got | %x != %x", ct, wantCT)
		}
	})

}

func TestDecrypt(t *testing.T) {
	ct, err := base64.StdEncoding.DecodeString("efrjHxns8qKwYtQJolQZdv6+HNRb1teO8xcxM1/3/LKJhsjvfYsXPOZLjXY=")
	if err != nil {
		t.Fatalf("failed to decode bin string")
	}
	msg := []byte("Hello World\n")

	t.Run("successful decryption", func(t *testing.T) {
		pt, err := Decrypt(ct, msg, msg)
		if err != nil {
			t.Fatalf("wanted err == nil, got err = %v", err)
		}

		wantPT := []byte("Hello World\n")
		if !bytes.Equal(pt, wantPT) {
			t.Fatalf("want != got | %x != %x", pt, wantPT)
		}
	})

	t.Run("bad decryption", func(t *testing.T) {
		_, err := Decrypt(ct, []byte("bad key"), msg)
		if err == nil {
			t.Fatalf("wanted err != nil, got err = %v", err)
		}
	})

}

func TestEncryptDecrypt(t *testing.T) {
	msg := []byte("hello there")
	key := []byte("key")
	IV := []byte("IV")

	ct, err := Encrypt(msg, key, IV)
	if err != nil {
		t.Fatalf("wanted err == nil | got: %v", err)
	}
	pt, err := Decrypt(ct, key, IV)
	if err != nil {
		t.Fatalf("wanted err == nil | got: %v", err)
	}

	if !bytes.Equal(msg, pt) {
		t.Fatalf("want != got | %x != %x", ct, pt)
	}
}
