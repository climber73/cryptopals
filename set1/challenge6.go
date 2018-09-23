package cryptopals

import (
	"encoding/base64"
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
	//chunks := splitIntoChunks(input, keySize)

	res = make([]byte, len(input))
	key = make([]byte, 3)
	return
}

func evaluateKeySize(input []byte) int {
	var bestKeySize int
	var diff float64
	minDiff := math.Inf(0)
	for keySize := 2; keySize <= 40; keySize++ {
		chunks := splitIntoChunks(input, keySize)
		diff = 0
		for i := 0; i < len(chunks); i += 2 {
			diff += float64(hammingDistance(chunks[i], chunks[i+1])) / float64(keySize)
		}
		diff = diff / float64(len(chunks)/2)
		if diff < minDiff {
			minDiff = diff
			bestKeySize = keySize
		}
	}
	return bestKeySize
}

func splitIntoChunks(input []byte, chunkSize int) (chunks [][]byte) {
	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			break
		}
		chunks = append(chunks, input[i:end])
	}
	if len(chunks)%2 != 0 {
		chunks = chunks[:len(chunks)-1]
	}
	return
}
