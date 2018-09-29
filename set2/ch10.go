package set2

import "github.com/climber73/criptopals/common"

func CBCDecrypt(siphertext []byte, key []byte, IV []byte) (text []byte) {
	size := len(key)
	text = make([]byte, len(siphertext))
	chunks := common.SplitIntoChunks(siphertext, size)
	var decrypted, prev, xored []byte
	for i := 0; i < len(chunks); i++ {
		if i == 0 {
			prev = IV
		} else {
			prev = chunks[i-1]
		}
		decrypted = common.ECBDecrypt(chunks[i], key)
		xored = common.Xor(decrypted, prev)
		copy(text[i*size:(i+1)*size], xored)
	}
	return
}
