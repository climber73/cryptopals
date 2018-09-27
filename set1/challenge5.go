package cryptopals

func RepeatingKeyXorApply(input []byte, key []byte) []byte {
	if len(key) == 0 {
		panic("Key length is 0")
	}
	res := make([]byte, len(input))
	var j int
	for i, b := range input {
		j = i % len(key)
		res[i] = b ^ key[j]
	}
	return res
}
