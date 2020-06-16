package toucan

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"
)

func TestEncrpt(t *testing.T) {
	msg := []byte("hi\n")
	data, key, IV := bytes.NewBuffer(msg), bytes.NewBuffer(msg), bytes.NewBuffer(msg)
	var out bytes.Buffer

	t.Run("successful encryption", func(t *testing.T) {
		wantBytes, err := base64.StdEncoding.DecodeString("Qe+lqxLDwVdK1jYbrEiHo238CCNJd7vWFO0Fl/5gw1jo1k/yONFJ0II36D2wfTbXbwJeneft4+5MUMlN/GBsHSvWYQ==")
		if err != nil {
			t.Fatalf("failed to decode base64 string")
		}

		if err := Encrypt(data, key, IV, &out); err != nil {
			t.Fatalf("encryption failed: %v", err)
		}

		if gotBytes := out.Bytes(); !bytes.Equal(wantBytes, gotBytes) {
			t.Fatalf("unexpected encryption result, %x != %x", wantBytes, gotBytes)
		}
	})

}

func TestDecrypt(t *testing.T) {
	wantBytes := []byte("hi\n")
	encryptedBytes, err := base64.StdEncoding.DecodeString("Qe+lqxLDwVdK1jYbrEiHo238CCNJd7vWFO0Fl/5gw1jo1k/yONFJ0II36D2wfTbXbwJeneft4+5MUMlN/GBsHSvWYQ==")
	if err != nil {
		t.Fatalf("failed to decode base64 string")
	}
	data, key, IV := bytes.NewBuffer(encryptedBytes), bytes.NewBuffer(wantBytes), bytes.NewBuffer(wantBytes)
	var out bytes.Buffer

	t.Run("successful decryption", func(t *testing.T) {

		if err := Decrypt(data, key, IV, &out); err != nil {
			t.Fatalf("decryption failed: %v", err)
		}

		if gotBytes := out.Bytes(); !bytes.Equal(wantBytes, gotBytes) {
			t.Fatalf("unexpected encryption result, %x != %x", wantBytes, gotBytes)
		}
	})

	data, key, IV = bytes.NewBuffer(encryptedBytes), bytes.NewBuffer(encryptedBytes), bytes.NewBuffer(encryptedBytes)
	out.Reset()

	t.Run("unsuccessful decryption", func(t *testing.T) {

		if err := Decrypt(data, key, IV, &out); err == nil {
			t.Fatalf("decryption succeeded, wanted failure")
		}

	})

}

func TestEncryptDecrypt(t *testing.T) {
	myMessage := "Hello Gopher!🤩"
	var encryptionResult bytes.Buffer
	var decryptionResult bytes.Buffer

	err := Encrypt(strings.NewReader(myMessage), strings.NewReader(myMessage), strings.NewReader(myMessage), &encryptionResult)
	if err != nil {
		t.Fatalf("encryption failed: %v", err)
	}

	err = Decrypt(&encryptionResult, strings.NewReader(myMessage), strings.NewReader(myMessage), &decryptionResult)
	if err != nil {
		t.Fatalf("decryption failed: %v", err)
	}

	if got := decryptionResult.String(); got != myMessage {
		t.Fatalf("Decrypt(Encrypt()) failed, got %s wanted %s", got, myMessage)
	}
}
