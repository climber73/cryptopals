package set1

import "testing"
import "encoding/base64"
import "io/ioutil"

func TestAes128EcbModeDecrypt(t *testing.T) {
	path := "test_data/7.txt"
	key := []byte("YELLOW SUBMARINE")
	text, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Can't read file: %s", path)
	}
	input, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		t.Fatalf("Can't decode base64 to bytes: %s", err)
	}
	res := Aes128EcbModeDecrypt(input, key)
	t.Logf("Result: %s", res)
}