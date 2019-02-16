package common

import "testing"
import "encoding/hex"

func TestECBDecrypt(t *testing.T) {
	var input, key []byte
	var want, res string
	input, _ = hex.DecodeString("f2f482b6fa95a89d671620cedfd8017e")
	key = []byte("YELLOW SUBMARINE")
	res = string(ECBDecrypt(input, key))
	want = "1234567890"
	check(res, want, t)

	input, _ = hex.DecodeString("93e463061d67c74aa6e5993aeaa4646860fa36707e45f499dba0f25b922301a5")
	key = []byte("YELLOW SUBMARINE")
	res = string(ECBDecrypt(input, key))
	want = "1234567890123456"
	check(res, want, t)

	input, _ = hex.DecodeString("93e463061d67c74aa6e5993aeaa464688cd081bea60d100b872eb6d5b75796d4")
	key = []byte("YELLOW SUBMARINE")
	res = string(ECBDecrypt(input, key))
	want = "12345678901234567"
	check(res, want, t)

	input, _ = hex.DecodeString("eadcc5aa4800dff175a49cf3a0f2041db9d8f63daefa0bb367819e289dae873ebcf95a6077f0bfe09385a722e577e026a987228246cf4dbcdc504262877a5694615cec9a624a79a84d20c964930cd9fddac0afab057057e60e01b983ff910d0bca11f45e799b602f05a08599fc2a1ed689ab0e3f14c7bad744f00988a793e38fecbfc8d7365c6c1fb67d5311c67754ec9a93fafb6aa2b71833b548f02b5ac2cf2d312e0c52f12ff9bb48bcc27c29b379b72318d2145e11be2f75d80bca19426f7ab9bce3f3f39dea8789a2f582dc7fe50f58d862b2483554dc05e9f3e906c797b76c8683e2db6adb6876fb662b8bdab24763178c0ccc682f02a137ab63fbe998")
	key = []byte("YELLOW SUBMARINE")
	res = string(ECBDecrypt(input, key))
	want = "Lorem ipsum dolor sit amet, eum quot consequat an, ut debitis prodesset complectitur per. Te paulo oporteat eos. Luptatum legendos explicari quo at. Quis dicam mea ne, te sea everti urbanitas, nisl omnes habemus vel ei. Utinam dictas consulatu vix ei."
	check(res, want, t)
}

func TestECBEncrypt(t *testing.T) {
	var input, key []byte
	var want, res string

	input = []byte("1234567890")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(ECBEncrypt(input, key))
	want = "f2f482b6fa95a89d671620cedfd8017e"
	check(res, want, t)

	input = []byte("1234567890123456")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(ECBEncrypt(input, key))
	want = "93e463061d67c74aa6e5993aeaa4646860fa36707e45f499dba0f25b922301a5"
	check(res, want, t)

	input = []byte("12345678901234567")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(ECBEncrypt(input, key))
	want = "93e463061d67c74aa6e5993aeaa464688cd081bea60d100b872eb6d5b75796d4"
	check(res, want, t)

	input = []byte("Lorem ipsum dolor sit amet, eum quot consequat an, ut debitis prodesset complectitur per. Te paulo oporteat eos. Luptatum legendos explicari quo at. Quis dicam mea ne, te sea everti urbanitas, nisl omnes habemus vel ei. Utinam dictas consulatu vix ei.")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(ECBEncrypt(input, key))
	want = "eadcc5aa4800dff175a49cf3a0f2041db9d8f63daefa0bb367819e289dae873ebcf95a6077f0bfe09385a722e577e026a987228246cf4dbcdc504262877a5694615cec9a624a79a84d20c964930cd9fddac0afab057057e60e01b983ff910d0bca11f45e799b602f05a08599fc2a1ed689ab0e3f14c7bad744f00988a793e38fecbfc8d7365c6c1fb67d5311c67754ec9a93fafb6aa2b71833b548f02b5ac2cf2d312e0c52f12ff9bb48bcc27c29b379b72318d2145e11be2f75d80bca19426f7ab9bce3f3f39dea8789a2f582dc7fe50f58d862b2483554dc05e9f3e906c797b76c8683e2db6adb6876fb662b8bdab24763178c0ccc682f02a137ab63fbe998"
	check(res, want, t)
}

func check(res, want string, t *testing.T) {
	if res != want {
		t.Errorf("\nExpected: '%x',\n     got: '%x'", want, res)
	}
}
