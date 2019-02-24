package set2

import "testing"
import "bytes"

func TestGuessMode(t *testing.T) {
	block := bytes.Repeat([]byte("a"), 16*7)
	encrypted, appliedMode := ProduceRandomlyEcryptedText(block)
	mode := GuessMode(encrypted)
	if appliedMode != mode {
		t.Errorf("expected:\n<%x>\ngot:\n<%x>", appliedMode, mode)
	}
}
