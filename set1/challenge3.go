package cryptopals

// http://cryptopals.com/sets/1/challenges/3

import "fmt"
import "encoding/hex"

var m = make(map[byte]byte)
var scores = make(map[byte]int)

func FindXorCipher(input string) byte {
	fmt.Printf("Input: %s\n", input)
	bytes, err := hex.DecodeString(input)
	if err != nil {
		panic("Can't decode input string")
	}
	tmp := make([]byte, len(bytes))
	for key := byte(0x00); key < 0xff; key++ {
		copy(tmp, bytes)
		initMap()
		decipherBytes(tmp, key)
		fmt.Printf("Key: %c ", key)
		scores[key] = evaluateScore(tmp)
	}

	max := 0
	var key byte
	for k, s := range scores {
		if s > max {
			max = s
			key = k
		}
	}
	return key
}

func initMap() {
	for c := 'a'; c < 'z'; c++ {
		m[byte(c)] = 0
	}
}

func evaluateScore(bytes []byte) int {
	fmt.Printf("Eval: \"%s\"... ", bytes)
	for c := byte('a'); c < byte('z'); c++ {
		for i := 0; i < len(bytes); i++ {
			if c == bytes[i] {
				m[c]++
			}
		}
	}
	sum := 0
	for c := range m {
		sum += int(m[c])
	}
	fmt.Printf("sum=%d\n", sum)
	return sum
}

func decipherBytes(bytes []byte, c byte) {
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i] ^ c
	}
}
