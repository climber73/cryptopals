package set1

import "bytes"
import "github.com/climber73/cryptopals/common"

func DetectAesECB(input []byte) bool {
	chunks := common.SplitIntoChunks(input, 16)
	for j := 0; j < len(chunks); j++ {
		for k := j + 1; k < len(chunks)-1; k++ {
			if j != k && bytes.Equal(chunks[j], chunks[k]) {
				return true
			}
		}
	}
	return false
}
