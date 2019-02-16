package common

import "bytes"

func PKCS7Pad(input []byte, size int) []byte {
	if size <= 1 || size >= 256 {
		panic("PKCS7Pad: wrong block size")
	}
	remainder := size - len(input)%size
	padding := []byte{byte(remainder)}
	return append(input, bytes.Repeat(padding, remainder)...)
}

func PKCS7Unpad(input []byte) []byte {
	paddingLen := int(input[len(input)-1])
	if paddingLen < 0 || paddingLen >= len(input) {
		panic("PKCS7Unpad: wrong padding len")
	}
	return input[:len(input)-paddingLen]
}
