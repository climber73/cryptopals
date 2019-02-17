package set2

import "testing"
import "io/ioutil"
import "encoding/base64"
import "bytes"

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

	decryptedPath := "test_data/10-decrypted.txt"
	decryptedText, err := ioutil.ReadFile(decryptedPath)
	if err != nil {
		t.Errorf("Can't read file: %s", decryptedPath)
	}
	if bytes.Compare(decryptedText, res) != 0 {
		t.Errorf("expected:\n<%x>\ngot:\n<%x>", decryptedText, res)
	}
}
