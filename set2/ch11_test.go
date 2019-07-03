package set2

import "testing"
import "bytes"
import "fmt"
import "gotest.tools/assert"

func TestGuessMode(t *testing.T) {
	block := bytes.Repeat([]byte("abcdefghijklmnop"), 3)
	for i := 0; i < 10; i++ {
		fmt.Printf("# %d:\n", i)
		encrypted, appliedMode := ProduceRandomlyEcryptedText(block)
		mode := GuessMode(encrypted)
		assert.Equal(t, appliedMode, mode)
	}
}
