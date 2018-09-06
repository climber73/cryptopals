package cryptopals

import "testing"

func TestBase16ToBase64(t *testing.T) {
	res, err := Base16ToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		t.Error(err)
	}
	exp := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if res != exp {
		t.Errorf("res is wrong:\nexpected:%s\nreceived:%s", exp, res)
	}
}
