package gobpcrypt_test

import (
	"testing"

	"github.com/elemc/gobpcrypt"
)

const (
	hello        = "Привет"
	expectedHash = "a03f848c68ea826702279a140a22cdd1"
)

func TestEncryptBytes(t *testing.T) {
	testB := []byte(hello)
	encB := gobpcrypt.EncryptBytes(testB, pwd)
	hashB := gobpcrypt.BytesToHex(encB)
	if hashB != expectedHash {
		t.Fatalf("Wrong EncryptBytes [%s]", hashB)
	}
}

func TestDecryptBytes(t *testing.T) {
	encB := gobpcrypt.HexToBytes(expectedHash)
	decB := gobpcrypt.DecryptBytes(encB, pwd)

	testB := string(decB)
	if testB != hello {
		t.Fatalf("Wrong DecryptBytes [%s]", testB)
	}
}

func TestEncryptStringToHash(t *testing.T) {
	hashB := gobpcrypt.EncryptStringToHash(hello, pwd)
	if hashB != expectedHash {
		t.Fatalf("Wrong EncryptStringToHash [%s]", hashB)
	}
}

func TestDecryptHash(t *testing.T) {
	testB := gobpcrypt.DecryptHash(expectedHash, pwd)
	if testB != hello {
		t.Fatalf("Wrong DecryptBytes [%s]", testB)
	}
}
