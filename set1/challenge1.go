package cryptopals

import "fmt"
import "errors"
import "encoding/hex"
import "log"

func Canonical(s string) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", decoded)
}

func Base16ToBase64(str string) (string, error) {
	fmt.Println(str)
	bytes, err := Base16ToBytes(str)
	if err != nil {
		return "", err
	}
	fmt.Println(bytes)
	fmt.Printf("%s", bytes)
	res, err := BytesToBase64(bytes)
	if err != nil {
		return "", err
	}
	fmt.Println(res)
	return res, nil
}

func Base16ToBytes(str string) ([]byte, error) {
	len := len(str)
	if len%2 != 0 {
		return nil, errors.New("length of str isn't even")
	}
	bytes := make([]byte, len/2)
	for i := 0; i < len/2; i++ {
		c1, ok := fromHexChar(str[i*2])
		if !ok {
			return nil, errors.New("wrong char: " + str[i*2])
		}
		c2, ok := fromHexChar(str[i*2+1])
		if !ok {
			return nil, errors.New("wrong char: " + str[i*2+1])
		}
		b := c1<<4 + c2
		bytes[i] = b
	}
	return bytes, nil
}

func fromHexChar(c byte) (byte, bool) {
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

func BytesToBase64(bytes []byte) (string, error) {
	return "", nil
}
