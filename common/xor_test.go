package common

import "testing"

func TestXor(t *testing.T) {
	res := Xor([]byte{0xFF, 0xFF}, []byte{0x0A, 0x0F})
	if len(res) != 2 {
		t.Fail()
	}
	if res[0] != byte(0xF5) || res[1] != byte(0xF0) {
		t.Fail()
	}
}
