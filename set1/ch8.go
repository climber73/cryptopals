package cryptopals

import "io/ioutil"
import "encoding/hex"
import "log"
import "strings"

func DetectAes(path string) {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't read file: %s", path)
	}
	found := -1
	for i, line := range strings.Split(string(text), "\n") {
		if len(line) == 0 {
			break
		}
		input, err := hex.DecodeString(line)
		if err != nil {
			log.Fatalf("Can't decode hex to bytes: %s", err)
		}
		chunks := splitIntoChunks(input, 16)
		for j := 0; j < len(chunks); j++ {
			for k := j + 1; k < len(chunks)-1; k++ {
				if j != k && hammingDistance(chunks[j], chunks[k]) == 0 {
					found = i
				}
			}
		}
	}
	log.Printf("Line: %d", found)
}
