package cryptopals

func PKCSPadBlock(block []byte, size int) (padded []byte) {
	diff := size - len(block)
	if diff < 0 {
		panic("Can't pad the block to a smaller size")
	}
	if diff > 255 {
		panic("Can't pad the block when diff is more than 255")
	}
	padded = make([]byte, size)
	copy(padded, block)
	if diff > 0 {
		for i := len(block); i < size; i++ {
			padded[i] = byte(diff)
		}
	}
	return
}
