package gobpcrypt_test

import (
	"testing"

	"github.com/elemc/gobpcrypt"
	"github.com/elemc/gobpcrypt/oldformat"
)

const (
	encOldHash123 = "0ffeae79788ab030f96a4aa30ec819e85e431d3a7e95cc18d2bf06a95bfcec74"
	encOldHash234 = "0ffeae79788ab030f96a4aa30ec819e8055c3092d71271fc1d0711ff033a609d"
	pwd           = "3510"
)

// oldformat
func TestOldDecryptBytes(t *testing.T) {
	bhash := gobpcrypt.HexToBytes(encOldHash123)
	dd := oldformat.OldDecryptBytes(bhash, pwd)
	if string(dd) != "123" {
		t.Fatal("Wrong OldDecryptBytes")
	}
}

func TestOldDecryptHash(t *testing.T) {
	result := oldformat.OldDecryptHash(encOldHash234, pwd)
	if result != "234" {
		t.Fatal("Wrong OldDecryptHash")
	}
}
