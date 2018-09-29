package common

func SplitIntoChunks(input []byte, chunkSize int) (chunks [][]byte) {
	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end >= len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}
	return
}

func TransposeChunks(chunks [][]byte) (transposed [][]byte) {
	n := len(chunks[0])
	for i := 0; i < n; i++ {
		row := make([]byte, len(chunks))
		for j := range chunks {
			row[j] = chunks[j][i]
		}
		transposed = append(transposed, row)
	}
	return
}
