package set2

import (
	"github.com/climber73/criptopals/common"
)

func PKCS7Pad(input []byte, size int) []byte {
	return common.PKCS7Pad(input, size)
}
