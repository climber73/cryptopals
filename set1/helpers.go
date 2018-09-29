package set1

import "io/ioutil"
import "encoding/hex"
import "log"

const ENGLISH_TEXT_FILE = "test_data/peterpan.txt"

type Corpus map[byte]float32

func buildCorpus() Corpus {
	text, err := ioutil.ReadFile(ENGLISH_TEXT_FILE)
	if err != nil {
		log.Panicf("can't read file: %s", ENGLISH_TEXT_FILE)
	}
	corpus := make(Corpus)
	for _, b := range text {
		corpus[b]++
	}
	length := float32(len(text))
	for k, v := range corpus {
		corpus[k] = v / length
	}
	return corpus
}

func decodeHex(s string) []byte {
	res, err := hex.DecodeString(s)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func hammingDistance(a, b []byte) int {
	if len(a) != len(b) {
		log.Fatalf("Mismatched lengths: len(%s)=%d, len(%s)=%d", a, len(a), b, len(b))
	}
	var diff int
	for i := range a {
		diff += int(countBits(a[i] ^ b[i]))
	}
	return diff
}

func countBits(b byte) (res byte) {
	for b > 0 {
		res += b & 0x01
		b >>= 1
	}
	return
}
