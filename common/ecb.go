package common

import "crypto/aes"

func ECBDecrypt(input []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	res := make([]byte, len(input))
	for i := 0; i < len(input); i += len(key) {
		block.Decrypt(res[i:i+len(key)], input[i:i+len(key)])
	}
	return res
}

func ECBEncrypt(input []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	res := make([]byte, len(input))
	for i := 0; i < len(input); i += len(key) {
		block.Encrypt(res[i:i+len(key)], input[i:i+len(key)])
	}
	return res
}
