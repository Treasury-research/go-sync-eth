package util

import "encoding/hex"

func DecodeStr(hex string) string {
	for i := len(hex); i > 0; i = i - 2 {
		s := hex[i-2 : i]
		if s != "00" {
			return hex[:i]
		}
	}
	return hex
}

// Hex2Bytes returns the bytes represented by the hexadecimal string str.
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
