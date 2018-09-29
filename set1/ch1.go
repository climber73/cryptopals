package cryptopals

// http://cryptopals.com/sets/1/challenges/1

import "fmt"
import "errors"

func Base16ToBase64(str string) (string, error) {
	bytes, err := base16ToBytes(str)
	if err != nil {
		return "", err
	}
	res, err := bytesToBase64(bytes)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func base16ToBytes(str string) ([]byte, error) {
	len := len(str)
	if len%2 != 0 {
		return nil, errors.New("length of str isn't even")
	}
	bytes := make([]byte, len/2)
	for i := 0; i < len/2; i++ {
		c1, ok := hexCharToByte(str[i*2])
		if !ok {
			return nil, errors.New(fmt.Sprintf("wrong char: %c", str[i*2]))
		}
		c2, ok := hexCharToByte(str[i*2+1])
		if !ok {
			return nil, errors.New(fmt.Sprintf("wrong char: %c", str[i*2+1]))
		}
		b := c1<<4 + c2
		bytes[i] = b
	}
	return bytes, nil
}

func bytesToBase64(bytes []byte) ([]byte, error) {
	n := (len(bytes) / 3) * 3
	rest := len(bytes) - n
	result := make([]byte, outpuSize(len(bytes)))
	var i, oi int
	var ok bool
	for i < n {
		result[oi+0], result[oi+1], result[oi+2], result[oi+3], ok =
			encodeThreeBytesToFourBase64Chars(bytes[i+0], bytes[i+1], bytes[i+2])
		if !ok {
			return result, errors.New(fmt.Sprintf("wrong bytes: %b, %b, %b", bytes[i+0], bytes[i+1], bytes[i+2]))
		}
		oi += 4
		i += 3
	}
	if rest == 1 {
		result[oi+0], result[oi+1], _, _, ok =
			encodeThreeBytesToFourBase64Chars(bytes[i+0], 0x00, 0x00)
		if !ok {
			return result, errors.New(fmt.Sprintf("wrong byte: %b", bytes[i+0]))
		}
		result[oi+2] = '='
		result[oi+3] = '='
	}
	if rest == 2 {
		result[oi+0], result[oi+1], result[oi+2], _, ok =
			encodeThreeBytesToFourBase64Chars(bytes[i+0], bytes[i+1], 0x00)
		if !ok {
			return result, errors.New(fmt.Sprintf("wrong bytes: %b, %b", bytes[i+0], bytes[i+1]))
		}
		result[oi+3] = '='
	}
	return result, nil
}

func encodeThreeBytesToFourBase64Chars(b1, b2, b3 byte) (byte, byte, byte, byte, bool) {
	val := uint(b1)<<16 | uint(b2)<<8 | uint(b3)
	a, aok := toBase64Char(byte(val & 0xFC0000 >> 18))
	b, bok := toBase64Char(byte(val & 0x03F000 >> 12))
	c, cok := toBase64Char(byte(val & 0x000FC0 >> 6))
	d, dok := toBase64Char(byte(val & 0x00003F))
	return a, b, c, d, aok && bok && dok && cok
}

func outpuSize(inputSize int) int {
	n := inputSize / 3
	r := inputSize % 3
	if r == 0 {
		return 4 * n
	} else {
		return 4 * (n + 1)
	}
}

func toBase64Char(b byte) (byte, bool) {
	switch {
	case 0 <= b && b <= 25:
		return b + 'A', true
	case 26 <= b && b <= 51:
		return b - 26 + 'a', true
	case 52 <= b && b <= 61:
		return b - 52 + '0', true
	case b == 62:
		return '+', true
	case b == 63:
		return '/', true
	}
	fmt.Printf("Error: wrong byte: %b\n", b)
	return 0, false
}

func hexCharToByte(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	}
	return 0, false
}
