package set1

import (
	"encoding/base64"
	"github.com/climber73/cryptopals/common"
	"io/ioutil"
	"log"
	"math"
)

func BreakRepeatingKeyXor(path string) (res []byte, key []byte) {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't read file: %s", path)
	}
	input, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		log.Fatalf("Can't decode base64 to bytes: %s", err)
	}

	keySize := evaluateKeySize(input)
	log.Printf("Best key size: %d", keySize)
	chunks := splitIntoNormalizedChunks(input, keySize)
	transposed := common.TransposeChunks(chunks)
	key = solveEachBlock(transposed)
	res = RepeatingKeyXorApply(input, key)
	return
}

func solveEachBlock(transposed [][]byte) (key []byte) {
	key = make([]byte, len(transposed))
	corpus := buildCorpus()
	for i := range transposed {
		_, key[i], _ = DecodeXoredBytes(transposed[i], corpus)
	}
	return
}

func evaluateKeySize(input []byte) int {
	var bestKeySize int
	var diff float64
	minDiff := math.Inf(0)
	for keySize := 2; keySize <= 40; keySize++ {
		chunks := splitIntoNormalizedChunks(input, keySize)
		diff = 0
		for i := 0; i < len(chunks); i += 2 {
			if len(chunks[i+1]) != keySize {
				break
			}
			hd, err := common.HammingDistance(chunks[i], chunks[i+1])
			if err != nil {
				panic(err.Error())
			}
			diff += float64(hd) / float64(keySize)
		}
		diff = diff / float64(len(chunks)/2)
		if diff < minDiff {
			minDiff = diff
			bestKeySize = keySize
		}
	}
	return bestKeySize
}

func splitIntoNormalizedChunks(input []byte, keySize int) [][]byte {
	chunks := common.SplitIntoChunks(input, keySize)
	if len(chunks[len(chunks)-1]) != keySize {
		chunks = chunks[:len(chunks)-1]
	}
	if len(chunks)%2 != 0 { // we need pairs
		chunks = chunks[:len(chunks)-1]
	}
	return chunks
}
