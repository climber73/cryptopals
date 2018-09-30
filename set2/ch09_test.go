package set2

import "testing"

func TestPKCS7Pad(t *testing.T) {
	var input, want, padded []byte

	input = []byte{7, 7, 7}
	padded = PKCS7Pad(input, 3)
	want = []byte{7, 7, 7, 3, 3, 3}
	checkResult(t, want, padded)

	input = []byte{7, 7, 7}
	padded = PKCS7Pad(input, 5)
	want = []byte{7, 7, 7, 2, 2}
	checkResult(t, want, padded)

	input = []byte{7, 7, 7, 7, 7, 7, 7}
	padded = PKCS7Pad(input, 3)
	want = []byte{7, 7, 7, 7, 7, 7, 7, 2, 2}
	checkResult(t, want, padded)
}

func checkResult(t *testing.T, want, padded []byte) {
	if len(padded) != len(want) {
		t.Errorf("expected len %d, got: %d", len(want), len(padded))
	}
	for i := 0; i < len(want); i++ {
		if padded[i] != want[i] {
			t.Errorf("expected: %d, got: %d", want, padded)
		}
	}
}
