package gobpcrypt

import (
	"fmt"
	"log"
	"strings"
)

// BytesToHex method convert byte slice to hex string
func BytesToHex(buffer []byte) string {
	var list []string
	for _, b := range buffer {
		list = append(list, fmt.Sprintf("%02x", b))
	}
	result := strings.Join(list, "")
	return result
}

// HexToBytes method convert hex string to byte slice
func HexToBytes(hex string) (result []byte) {
	for i := 0; i < len(hex); i = i + 2 {
		var value byte
		_, err := fmt.Sscanf(hex[i:i+2], "%02x", &value)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, value)
	}
	return
}
