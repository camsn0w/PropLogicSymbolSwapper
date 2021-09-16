package main

import (
	"errors"
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
			err = popnext(data, i)
			if err != nil {
				return err
			}
		case '~', '!':
			(*data)[i] = '¬'
		case 'n':
			(*data)[i] = '¬'
			err = popnext(data, i)
			if err != nil {
				return err
			}
			err = popnext(data, i)
			if err != nil {
				return err
			}
		case 'o':
			(*data)[i] = '∨'
		case 'a':
			(*data)[i] = '∧'
			err = popnext(data, i)
			if err != nil {
				return err
			}
			err = popnext(data, i)
			if err != nil {
				return err
			}
		default:
			continue
		}
	}
	return nil
}

func popnext(data *[]rune, index int) error {
	if index == len(*data)-1 {
		return errors.New("Invalid formatting at index " + string(rune(index)))
	}

	*data = append((*data)[:index+1], (*data)[index+2:]...)
	return nil
}
