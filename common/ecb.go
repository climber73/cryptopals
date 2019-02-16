package common

import "crypto/aes"

func ECBDecrypt(input []byte, key []byte) []byte {
	res := ECBDecryptPadded(input, key)
	return PKCS7Unpad(res)
}

func ECBDecryptPadded(input []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	res := make([]byte, len(input))
	for i := 0; i < len(input); i += len(key) {
		end := i + len(key)
		if end >= len(input) {
			end = len(input)
		}
		block.Decrypt(res[i:end], input[i:end])
	}
	return res
}

func ECBEncrypt(input []byte, key []byte) []byte {
	padded := PKCS7Pad(input, len(key))
	return ECBEncryptPadded(padded, key)
}

func ECBEncryptPadded(input []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	res := make([]byte, len(input))
	for i := 0; i < len(input); i += len(key) {
		end := i + len(key)
		if end >= len(input) {
			end = len(input)
		}
		block.Encrypt(res[i:end], input[i:end])
	}
	return res
}
