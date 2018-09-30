package set2

import "testing"

func TestGuessMode(t *testing.T) {
	block := []byte("1234567890")
	GuessMode(block)
}
