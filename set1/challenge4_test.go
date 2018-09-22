package cryptopals

import "testing"

func TestFindXoredString(t *testing.T) {
	path := "test_data/4.txt"
	res, key := FindXoredString(path)
	t.Logf("Decoded string: \"%s\" (key: %c)", res, key)
}
