package common

func Xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("lengths are different")
	}
	res := make([]byte, len(a))
	for i := range res {
		res[i] = a[i] ^ b[i]
	}
	return res
}
