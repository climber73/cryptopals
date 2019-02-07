package set1

import "fmt"
import "encoding/base64"
import "encoding/hex"
import "io/ioutil"
import "github.com/climber73/cryptopals/common"

func AesInEcbModeDecrypt(path string, key []byte) ([]byte, error) {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Can't read file: %s", path)
	}
	input, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, fmt.Errorf("Can't decode base64 to bytes: %s", err)
	}
	fmt.Printf("%s", hex.EncodeToString(input[:32]))
	res := common.ECBDecrypt(input, key)
	return res, nil
}
