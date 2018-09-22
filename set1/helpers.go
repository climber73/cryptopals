package cryptopals

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
