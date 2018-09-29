package cryptopals

import "testing"

func TestFixedXor(t *testing.T) {
	res := FixedXor([]byte{0xFF, 0xFF}, []byte{0x0A, 0x0F})
	if len(res) != 2 {
		t.Fail()
	}
	if res[0] != byte(0xF5) || res[1] != byte(0xF0) {
		t.Fail()
	}
}
