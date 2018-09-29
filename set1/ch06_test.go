package set1

import "testing"

func TestBreakRepeatingKeyXor(t *testing.T) {
	inputFilePath := "test_data/6.txt"
	res, key := BreakRepeatingKeyXor(inputFilePath)
	t.Logf("Result: %s (key: %s)", res, key)
}
