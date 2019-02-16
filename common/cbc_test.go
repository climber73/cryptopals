package common

import "testing"
import "encoding/hex"

func TestCBCDecrypt(t *testing.T) {
	var input, key, IV []byte
	var want, res string

	IV = make([]byte, 16)
	input, _ = hex.DecodeString("f2f482b6fa95a89d671620cedfd8017e")
	key = []byte("YELLOW SUBMARINE")
	res = string(CBCDecrypt(input, key, IV))
	want = "1234567890"
	check(res, want, t)

	IV = make([]byte, 16)
	input, _ = hex.DecodeString("93e463061d67c74aa6e5993aeaa46468d06ad8c3f98cb2285cc7c19bbd2fbe39")
	key = []byte("YELLOW SUBMARINE")
	res = string(CBCDecrypt(input, key, IV))
	want = "1234567890123456"
	check(res, want, t)

	IV = make([]byte, 16)
	input, _ = hex.DecodeString("93e463061d67c74aa6e5993aeaa464685f1e054143ae370e679be5bcbaabe133")
	key = []byte("YELLOW SUBMARINE")
	res = string(CBCDecrypt(input, key, IV))
	want = "12345678901234567"
	check(res, want, t)

	IV = make([]byte, 16)
	input, _ = hex.DecodeString("eadcc5aa4800dff175a49cf3a0f2041d9e12e8119a2f31fd3b4059a5f8af255404090299ced83e197692a75e342ad433f202978bf08ae3fadbdaa0313d0ae086143cab6182ebae28265a9162838cb0d1d4a6c53c0f7bd6013830bc3bd5572f879d09eabfe9d72a993bf0bb4a9f2242167bd97fa5e55f72c092a983b1d3941c91cf3d949cc200e7d3548e5b0b4d5fe605b1b4855287870e52e8f0301e826b018fede06f331d5ae0789da818be19623a46ed38f51fb2b1c75c7290035536ee930d60d739f00e289757dda82d0ebac3186cc8f32e879a5368d52e3b4ddd80b5d64e45e8fb07bb107da8d249559e8af62c351c3d9a4d7ede523670c0528f6debe2ee")
	key = []byte("YELLOW SUBMARINE")
	res = string(CBCDecrypt(input, key, IV))
	want = "Lorem ipsum dolor sit amet, eum quot consequat an, ut debitis prodesset complectitur per. Te paulo oporteat eos. Luptatum legendos explicari quo at. Quis dicam mea ne, te sea everti urbanitas, nisl omnes habemus vel ei. Utinam dictas consulatu vix ei."
	check(res, want, t)
}

func TestCBCEncrypt(t *testing.T) {
	var input, key, IV []byte
	var want, res string

	IV = make([]byte, 16)
	input = []byte("1234567890")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(CBCEncrypt(input, key, IV))
	want = "f2f482b6fa95a89d671620cedfd8017e"
	check(res, want, t)

	IV = make([]byte, 16)
	input = []byte("1234567890123456")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(CBCEncrypt(input, key, IV))
	want = "93e463061d67c74aa6e5993aeaa46468d06ad8c3f98cb2285cc7c19bbd2fbe39"
	check(res, want, t)

	IV = make([]byte, 16)
	input = []byte("12345678901234567")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(CBCEncrypt(input, key, IV))
	want = "93e463061d67c74aa6e5993aeaa464685f1e054143ae370e679be5bcbaabe133"
	check(res, want, t)

	IV = make([]byte, 16)
	input = []byte("Lorem ipsum dolor sit amet, eum quot consequat an, ut debitis prodesset complectitur per. Te paulo oporteat eos. Luptatum legendos explicari quo at. Quis dicam mea ne, te sea everti urbanitas, nisl omnes habemus vel ei. Utinam dictas consulatu vix ei.")
	key = []byte("YELLOW SUBMARINE")
	res = hex.EncodeToString(CBCEncrypt(input, key, IV))
	want = "eadcc5aa4800dff175a49cf3a0f2041d9e12e8119a2f31fd3b4059a5f8af255404090299ced83e197692a75e342ad433f202978bf08ae3fadbdaa0313d0ae086143cab6182ebae28265a9162838cb0d1d4a6c53c0f7bd6013830bc3bd5572f879d09eabfe9d72a993bf0bb4a9f2242167bd97fa5e55f72c092a983b1d3941c91cf3d949cc200e7d3548e5b0b4d5fe605b1b4855287870e52e8f0301e826b018fede06f331d5ae0789da818be19623a46ed38f51fb2b1c75c7290035536ee930d60d739f00e289757dda82d0ebac3186cc8f32e879a5368d52e3b4ddd80b5d64e45e8fb07bb107da8d249559e8af62c351c3d9a4d7ede523670c0528f6debe2ee"
	check(res, want, t)
}
