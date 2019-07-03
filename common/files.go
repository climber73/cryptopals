package common

import "io/ioutil"

func Load(path string) []byte {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return text
}
