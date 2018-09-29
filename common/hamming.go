package common

import "fmt"

func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Mismatched lengths: len(%d)=%d, len(%d)=%d", a, len(a), b, len(b))
	}
	var diff int
	for i := range a {
		diff += int(countBits(a[i] ^ b[i]))
	}
	return diff, nil
}

func countBits(b byte) (res byte) {
	for b > 0 {
		res += b & 0x01
		b >>= 1
	}
	return
}
