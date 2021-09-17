package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		print("Please run with a filename as the argument")
		os.Exit(-1)
	}

	filename := os.Args[1]
	data, err := ReadFile(filename)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	err = swap(&data)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

	err = WriteFile(filename, data)
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}

}

func swap(data *[]rune) error {
	var err error
	for i := 0; i < len(*data); i++ {
		switch (*data)[i] {
		case 'v':
			(*data)[i] = '∨'
		case '&', '^':
			(*data)[i] = '∧'
		case '=':
			(*data)[i] = '⇒'
			err = remove(data, i+1, 1)
		case '~', '!':
			(*data)[i] = '¬'
		case 'n':
			(*data)[i] = '¬'
			err = remove(data, i+1, 2)
		case 'o':
			(*data)[i] = '∨'
			err = remove(data, i+1, 1)
		case 'a':
			(*data)[i] = '∧'
			err = remove(data, i+1, 2)
		default:
			continue
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func remove(data *[]rune, off int, n int) error {
	if off+n >= len(*data) {
		return fmt.Errorf("invalid formatting at index: %v", off)
	}

	*data = append((*data)[:off], (*data)[off+n:]...)
	return nil
}
