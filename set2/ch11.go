package set2

import "math/rand"
import "fmt"
import "time"

func GuessMode(text []byte) {
	fmt.Printf("text: %d\n", text)
	rand.Seed(time.Now().UnixNano())
	key := generateRandomAesKey()
	fmt.Printf("key: %d\n", key)
	modified := modify(text)
	fmt.Printf("modified: %d\n", modified)

	//encrypted := make([]byte, len(modified))
	/*switch rand.Intn(2) {
	case 0:

	case 1:

	}*/
}

func chooseEncryption() {
}

func modify(text []byte) (modified []byte) {
	before := make([]byte, rand.Intn(6)+5)
	after := make([]byte, rand.Intn(6)+5)
	for i := range before {
		before[i] = byte(rand.Intn(256))
	}
	for i := range after {
		after[i] = byte(rand.Intn(256))
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
		key[i] = byte(rand.Intn(256))
	}
	return
}
