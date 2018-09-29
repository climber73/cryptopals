package set1

import "testing"

func TestAesInEcbModeDecrypt(t *testing.T) {
	path := "test_data/7.txt"
	key := []byte("YELLOW SUBMARINE")
	res, err := AesInEcbModeDecrypt(path, key)
	if err != nil {
		t.Fail()
	}
	t.Logf("Result: %s", res)
}
