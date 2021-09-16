package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]rune, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.Runes(contents), nil
}

func WriteFile(filename string, data []rune) error {
	f, err := os.Create("./swapped-" + filename)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(data))

	return err
}
