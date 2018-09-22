package cryptopals

import "testing"
import "fmt"
import "io/ioutil"
import "encoding/hex"

const ENGLISH_TEXT_FILE = "test_data/peterpan.txt"

func TestFindXorCipher(t *testing.T) {
	text, err := ioutil.ReadFile(ENGLISH_TEXT_FILE)
	if err != nil {
		t.Errorf("can't read file: %s", ENGLISH_TEXT_FILE)
	}
	corpus := buildCorpus(text)
	input := decodeHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", t)
	res, key := DecodeXoredBytes(input, corpus)
	fmt.Printf("%s (key: %c)\n", res, key)
}

func buildCorpus(text []byte) map[byte]float32 {
	corpus := make(map[byte]float32)
	for _, b := range text {
		corpus[b]++
	}
	length := float32(len(text))
	for k, v := range corpus {
		corpus[k] = v / length
	}
	return corpus
}

func decodeHex(s string, t *testing.T) []byte {
	res, err := hex.DecodeString(s)
	if err != nil {
		t.Error(err)
	}
	return res
}
