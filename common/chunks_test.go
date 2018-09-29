package common

import "testing"

func TestSplitIntoChunks(t *testing.T) {
	res := SplitIntoChunks([]byte("1234567890"), 3)
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
	res := TransposeChunks(chunks)
	want := [][]byte{{'1', '4'}, {'2', '5'}, {'3', '6'}}
	for i := range res {
		for j := range res[i] {
			if res[i][j] != want[i][j] {
				t.Fail()
			}
		}
	}
}
