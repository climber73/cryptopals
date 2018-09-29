package cryptopals

import "testing"

func TestCountBits(t *testing.T) {
	checkTestCountBitsResult(countBits(byte(0x00)), 0, t)
	checkTestCountBitsResult(countBits(byte(0xAA)), 4, t)
	checkTestCountBitsResult(countBits(byte(0xAF)), 6, t)
}

func checkTestCountBitsResult(res, exp byte, t *testing.T) {
	if res != exp {
		t.Fatalf("wrong count. Expected %d, got %d", exp, res)
	}
}

func TestHammingDistance(t *testing.T) {
	a := "this is a test"
	b := "wokka wokka!!!"
	res := hammingDistance([]byte(a), []byte(b))
	if res != 37 {
		t.Fatalf("wrong distance. Expected %d, got %d", 37, res)
	}
}

func TestSplitIntoChunks(t *testing.T) {
	res := splitIntoChunks([]byte("1234567890"), 3)
	want := [][]byte{{'1', '2', '3'}, {'4', '5', '6'}}
	for i := range res {
		for j := range res[i] {
			if res[i][j] != want[i][j] {
				t.Fail()
			}
		}
	}
}

func TestTransposeChunks(t *testing.T) {
	chunks := [][]byte{{'1', '2', '3'}, {'4', '5', '6'}}
	res := transposeChunks(chunks)
	want := [][]byte{{'1', '4'}, {'2', '5'}, {'3', '6'}}
	for i := range res {
		for j := range res[i] {
			if res[i][j] != want[i][j] {
				t.Fail()
			}
		}
	}
}

func TestBreakRepeatingKeyXor(t *testing.T) {
	inputFilePath := "test_data/6.txt"
	res, key := BreakRepeatingKeyXor(inputFilePath)
	t.Logf("Result: %s (key: %s)", res, key)
}
