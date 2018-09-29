package common

import "testing"
import "encoding/hex"

func TestECBDecrypt(t *testing.T) {
	var input, key []byte
	var want, res string
	input, _ = hex.DecodeString("091230aade3eb330dbaa4358f88d2a6c")
	key = []byte("YELLOW SUBMARINE")
	res = string(ECBDecrypt(input, key))
	want = "I'm back and I'm"
	if res != want {
		t.Errorf("Expected: %s, got: %s", want, res)
	}

	input, _ = hex.DecodeString("37b72d0cf4c22c344aec4142d00ce530")
	key = []byte("YELLOW SUBMARINE")
	res = string(ECBDecrypt(input, key))
	want = " ringin' the bel"
	if res != want {
		t.Errorf("Expected: %s, got: %s", want, res)
	}
}

func TestECBEncrypt(t *testing.T) {
	var input, key []byte
	var want, res string
	input = []byte("I'm back and I'm")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(ECBEncrypt(input, key))
	want = "091230aade3eb330dbaa4358f88d2a6c"
	if res != want {
		t.Errorf("Expected: %s, got: %s", want, res)
	}

	input = []byte(" ringin' the bel")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(ECBEncrypt(input, key))
	want = "37b72d0cf4c22c344aec4142d00ce530"
	if res != want {
		t.Errorf("Expected: %s, got: %s", want, res)
	}
}
