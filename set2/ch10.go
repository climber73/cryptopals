package set2

import "github.com/climber73/criptopals/common"

func CBCDecrypt(siphertext, key, IV []byte) []byte {
	return common.CBCDecrypt(siphertext, key, IV)
}
