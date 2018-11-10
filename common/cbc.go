package common

func CBCDecrypt(input, key, IV []byte) []byte {
	return cbcAction(Decrypt, input, key, IV)
}

func CBCEncrypt(input, key, IV []byte) []byte {
	return cbcAction(Encrypt, input, key, IV)
}

func cbcAction(action CryptoAction, input, key, IV []byte) []byte {
	size := len(key)
	output := make([]byte, len(input))
	chunks := SplitIntoChunks(input, size)
	var processed, prev, xored []byte
	for i := 0; i < len(chunks); i++ {
		if i == 0 {
			prev = IV
		} else {
			prev = chunks[i-1]
		}
		switch action {
		case Encrypt:
			xored = Xor(chunks[i], prev)
			processed = ECBEncrypt(xored, key)
			copy(output[i*size:(i+1)*size], processed)
		case Decrypt:
			processed = ECBDecrypt(chunks[i], key)
			xored = Xor(processed, prev)
			copy(output[i*size:(i+1)*size], xored)
		}
	}
	return output
}
