package set1

import "testing"

func TestFindXorCipher(t *testing.T) {
	corpus := buildCorpus()
	input := decodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	res, key, _ := DecodeXoredBytes(input, corpus)
	t.Logf("Decoded string: \"%s\" (key: %c)", res, key)
}
