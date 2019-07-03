package set1

import "testing"
import "strings"
import "log"
import "github.com/climber73/cryptopals/common"

func TestDetectAesECB(t *testing.T) {
	found := -1
	text := common.Load("test_data/8.txt")
	for i, line := range strings.Split(string(text), "\n") {
		if len(line) == 0 {
			break
		}
		input := decodeHex(line)
		if DetectAesECB(input) {
			found = i
		}
	}
	log.Printf("Line: %d", found)
}
