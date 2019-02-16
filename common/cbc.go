package common

func CBCDecrypt(input, key, IV []byte) []byte {
	size := len(key)
	res := make([]byte, len(input))
	chunks := SplitIntoChunks(input, size)
	var processed, prev, xored []byte
	for i := 0; i < len(chunks); i++ {
		if i == 0 {
			prev = IV
		} else {
			prev = chunks[i-1]
		}
		processed = ECBDecryptPadded(chunks[i], key)
		xored = Xor(processed, prev)
		copy(res[i*size:(i+1)*size], xored)
	}
	return PKCS7Unpad(res)
}

func CBCEncrypt(input, key, IV []byte) []byte {
	padded := PKCS7Pad(input, len(key))
	size := len(key)
	res := make([]byte, len(padded))
	chunks := SplitIntoChunks(padded, size)
	var processed, prev, xored []byte
	for i := 0; i < len(chunks); i++ {
		if i == 0 {
			prev = IV
		} else {
			prev = res[(i-1)*size : i*size]
		}
		xored = Xor(chunks[i], prev)
		processed = ECBEncryptPadded(xored, key)
		copy(res[i*size:(i+1)*size], processed)
	}
	return res
}
