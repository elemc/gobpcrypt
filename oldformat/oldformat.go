package oldformat

/*
#cgo CFLAGS: -I/home/alex/workspace/code/libbpcrypt
#cgo LDFLAGS: -L/home/alex/workspace/code/libbpcrypt/build -lbpcrypt
#include <bpcrypt/old_crypt.h>
#include <bpcrypt/misc.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/elemc/gobpcrypt"
)

// OldDecryptBytes method decrypt by old mechanism bytes slice
func OldDecryptBytes(data []byte, password string) (decryptData []byte) {
	var dbufSize C.bp_size_t
	cPwd := C.CString(password)
	cBuf := C.CString(string(data))

	decBuf := C.old_decrypt_buffer(C.CString_convert_to_BP_buffer(cBuf), C.bp_size_t(len(data)), C.CString_convert_to_BP_buffer(cPwd), C.bp_size_t(4), &dbufSize)

	C.free(unsafe.Pointer(cPwd))
	C.free(unsafe.Pointer(cBuf))

	rightDbuf := unsafe.Pointer(decBuf)
	decryptData = C.GoBytes(rightDbuf, C.int(dbufSize))
	C.free(rightDbuf)

	return
}

// OldDecryptHash methdo decrypt by old mech hash
func OldDecryptHash(hash string, password string) (decryptData string) {
	bhash := gobpcrypt.HexToBytes(hash)
	dd := OldDecryptBytes(bhash, password)
	return string(dd)
}
