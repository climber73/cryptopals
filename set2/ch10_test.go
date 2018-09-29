package set2

import "testing"
import "io/ioutil"
import "encoding/base64"

func TestCBCDecrypt(t *testing.T) {
	path := "test_data/10.txt"
	text, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Can't read file: %s", path)
	}
	siphertext, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		t.Errorf("Can't decode base64 to bytes: %s", err)
	}
	res := CBCDecrypt(siphertext, []byte("YELLOW SUBMARINE"), make([]byte, 16))
	t.Logf("Decrypted text:\n%s\n", res)
}
