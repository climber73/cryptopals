package set1

import (
	"bufio"
	"log"
	"os"
)

func FindXoredString(path string) (result []byte, key byte) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	corpus := buildCorpus()
	reader := bufio.NewReader(file)
	var line string
	var max_score float32
	result = make([]byte, 30)
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		bytes := decodeHex(line[:len(line)-1])
		decoded, k, score := DecodeXoredBytes(bytes, corpus)
		if score > max_score {
			max_score = score
			key = k
			copy(result, decoded)
		}
	}
	return
}
