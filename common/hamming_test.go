package common

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
	res, err := HammingDistance([]byte(a), []byte(b))
	if err != nil {
		t.Fail()
	}
	if res != 37 {
		t.Fatalf("wrong distance. Expected %d, got %d", 37, res)
	}
}
