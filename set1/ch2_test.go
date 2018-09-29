package set1

import "testing"

func TestFixedXor(t *testing.T) {
	res, err := FixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	}
	want := "746865206b696420646f6e277420706c6179"
	if res != want {
		t.Errorf("Expected: %s, actual: %s\n", want, res)
	}
}
