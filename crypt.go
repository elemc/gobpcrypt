package gobpcrypt

/*
#cgo CFLAGS: -I/home/alex/workspace/code/libbpcrypt
#cgo LDFLAGS: -L/home/alex/workspace/code/libbpcrypt/build -lbpcrypt
#include <bpcrypt/crypt.h>
#include <bpcrypt/misc.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// DecryptBytes method decrypt by new mechanism bytes slice
func DecryptBytes(data []byte, password string) (decryptData []byte) {
	var dbufSize C.bp_size_t
	cPwd := C.CString(password)
	cBuf := C.CString(string(data))

	decBuf := C.decrypt_buffer(C.CString_convert_to_BP_buffer(cBuf), C.bp_size_t(len(data)), C.CString_convert_to_BP_buffer(cPwd), C.bp_size_t(4), &dbufSize)

	C.free(unsafe.Pointer(cPwd))
	C.free(unsafe.Pointer(cBuf))

	rightDbuf := unsafe.Pointer(decBuf)
	decryptData = C.GoBytes(rightDbuf, C.int(dbufSize))
	C.free(rightDbuf)

	return
}

// EncryptBytes method decrypt by new mechanism bytes slice
func EncryptBytes(data []byte, password string) (encryptData []byte) {
	var dbufSize C.bp_size_t
	cPwd := C.CString(password)
	cBuf := C.CString(string(data))

	decBuf := C.encrypt_buffer(C.CString_convert_to_BP_buffer(cBuf), C.bp_size_t(len(data)), C.CString_convert_to_BP_buffer(cPwd), C.bp_size_t(4), &dbufSize)

	C.free(unsafe.Pointer(cPwd))
	C.free(unsafe.Pointer(cBuf))

	rightDbuf := unsafe.Pointer(decBuf)
	encryptData = C.GoBytes(rightDbuf, C.int(dbufSize))
	C.free(rightDbuf)

	return
}

// DecryptHash methdo decrypt by new mech hash
func DecryptHash(hash string, password string) (decryptData string) {
	bhash := HexToBytes(hash)
	dd := DecryptBytes(bhash, password)
	return string(dd)
}

// EncryptStringToHash method return hash encrypted bytes for string
func EncryptStringToHash(dataStr string, password string) (encryptedData string) {
	data := []byte(dataStr)
	encData := EncryptBytes(data, password)
	encryptedData = BytesToHex(encData)
	return
}
