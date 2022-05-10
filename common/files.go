package common

import (
	"errors"
	"os"
)

func ReadFile(file string) (string, error) {
	if file == "" {
		return "", errors.New("please provide a file")
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func StoreString(file string, data string) error {
	return StoreBytes(file, []byte(data))
}

func StoreBytes(file string, data []byte) error {
	if file == "" {
		return errors.New("please provide a file")
	}
	return os.WriteFile(file, data, 0644)
}