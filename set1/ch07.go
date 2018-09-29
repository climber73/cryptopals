package set1

import "crypto/aes"

func Aes128EcbModeDecrypt(input []byte, key []byte) []byte {
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
