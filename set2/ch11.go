package set2

import "github.com/climber73/cryptopals/common"
import "math/rand"
import "bytes"
import "time"
import "fmt"

type Mode int

const (
	ECB Mode = 0
	CBC Mode = 1
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GuessMode(encrypted []byte) Mode {
	chunks := common.SplitIntoChunks(encrypted, 16)
	if len(chunks) == 0 {
		panic("Empty chunks!")
	}
	last := chunks[len(chunks)-1]
	if len(last) != 16 {
		panic("len(encrypted) mod 16 != 0")
	}
	for i, ch1 := range chunks {
		for j, ch2 := range chunks {
			if i != j && bytes.Equal(ch1, ch2) {
				return ECB
			}
		}
	}
	return CBC
}

// return applied mode as second parameter for test purposes:
func ProduceRandomlyEcryptedText(text []byte) ([]byte, Mode) {
	key := generateRandomAesKey()
	modified := appendPrependRandomBytes(text)
	var encrypted []byte
	mode := Mode(rand.Intn(2))
	fmt.Printf("Mode selected: %x\n", mode)
	switch mode {
	case ECB:
		encrypted = common.ECBEncrypt(modified, key)
	case CBC:
		encrypted = common.CBCEncrypt(modified, key, key)
	}
	return encrypted, mode
}

func appendPrependRandomBytes(text []byte) (modified []byte) {
	before := make([]byte, rand.Intn(6)+5)
	after := make([]byte, rand.Intn(6)+5)
	for i := range before {
		before[i] = randByte()
	}
	for i := range after {
		after[i] = randByte()
	}
	modified = make([]byte, len(before)+len(text)+len(after))
	copy(modified[:len(before)], before)
	copy(modified[len(before):len(before)+len(text)], text)
	copy(modified[len(before)+len(text):], after)
	return
}

func generateRandomAesKey() (key []byte) {
	key = make([]byte, 16)
	for i := range key {
		key[i] = randByte()
	}
	return
}

func randByte() byte {
	return byte(rand.Intn(256))
}
