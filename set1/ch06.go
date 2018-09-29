package set1

import (
	"encoding/base64"
	"github.com/climber73/criptopals/common"
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
	chunks := common.SplitIntoChunks(input, keySize)
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
		chunks := common.SplitIntoChunks(input, keySize)
		diff = 0
		for i := 0; i < len(chunks); i += 2 {
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
