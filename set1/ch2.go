package set1

import "encoding/hex"
import "github.com/climber73/criptopals/common"

func FixedXor(hex1, hex2 string) (string, error) {
	a, err := hex.DecodeString(hex1)
	if err != nil {
		return "", err
	}
	b, err := hex.DecodeString(hex2)
	if err != nil {
		return "", err
	}
	res := common.Xor(a, b)
	return hex.EncodeToString(res), nil
}
