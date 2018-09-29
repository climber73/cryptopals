package cryptopals

// http://cryptopals.com/sets/1/challenges/3

func DecodeXoredBytes(input []byte, c Corpus) ([]byte, byte, float32) {
	decoded := make([]byte, len(input))
	tmp := make([]byte, len(input))
	var key byte
	var max_score float32
	for k := 0x00; k <= 0xff; k++ {
		decodeBytes(input, tmp, byte(k))
		score := evaluateScore(tmp, c)
		if score > max_score {
			key = byte(k)
			copy(decoded, tmp)
			max_score = score
		}
	}

	return decoded, key, max_score
}

func evaluateScore(bytes []byte, c Corpus) float32 {
	var score float32
	for _, b := range bytes {
		score += c[b]
	}
	return score
}

func decodeBytes(input []byte, output []byte, k byte) {
	if len(input) != len(output) {
		panic("different lengths")
	}
	for i, value := range input {
		output[i] = value ^ k
	}
}
