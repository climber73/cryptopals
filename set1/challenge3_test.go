package cryptopals

import "testing"
import "fmt"

func TestFindXorCipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	key := FindXorCipher(input)
	fmt.Printf("The key is '%c' (%d)\n", key, key)
}
