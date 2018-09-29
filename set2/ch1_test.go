package cryptopals

import "testing"

func TestPKCSPadBlock(t *testing.T) {
	var input, padded []byte

	input = []byte{7, 7, 7}
	padded = PKCSPadBlock(input, 3)
	checkResult(t, input, padded, 3)

	input = []byte{7, 7, 7}
	padded = PKCSPadBlock(input, 5)
	checkResult(t, input, padded, 5)
}

func checkResult(t *testing.T, input, padded []byte, size int) {
	if len(padded) != size {
		t.Errorf("wrong len: %d", len(padded))
	}
	for i := 0; i < len(input); i++ {
		if padded[i] != input[i] {
			t.Errorf("wrong value at %d position of paddded slice: %d, input: %d", i, padded, input)
		}
	}
	for i := len(input); i < len(padded); i++ {
		if padded[i] != byte(size-len(input)) {
			t.Errorf("wrong value at %d position of paddded slice: %d, input: %d", i, padded, input)
		}
	}
}
