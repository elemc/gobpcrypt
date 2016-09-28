package gobpcrypt_test

import (
	"bytes"
	"testing"

	"github.com/elemc/gobpcrypt"
)

func TestBytesToHex(t *testing.T) {
	testB := []byte{01, 02, 03, 10}
	hex := gobpcrypt.BytesToHex(testB)
	if hex != "0102030a" {
		t.Fatalf("Wrong BytesToHex [%s]", hex)
	}
}

func TestHexToBytes(t *testing.T) {
	testHex := "0102030a"
	expected := []byte{01, 02, 03, 10}

	testB := gobpcrypt.HexToBytes(testHex)

	if !bytes.Equal(testB, expected) {
		t.Fatalf("Wrong HexToBytes [%+v]", testB)
	}
}
