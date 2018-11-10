package common

import (
	"crypto/aes"
)

func ECBDecrypt(input []byte, key []byte) []byte {
	return ecbAction(Decrypt, input, key)
}

func ECBEncrypt(input []byte, key []byte) []byte {
	return ecbAction(Encrypt, input, key)
}

func ecbAction(action CryptoAction, input []byte, key []byte) []byte {
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
		switch action {
		case Encrypt:
			block.Encrypt(res[i:end], input[i:end])
		case Decrypt:
			block.Decrypt(res[i:end], input[i:end])
		}
	}
	return res
}