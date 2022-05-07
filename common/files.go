package common

import (
	"errors"
	"os"
)

func ReadFile(file string) (string, error) {
	if file == "" {
		return "", errors.New("Please provide a file")
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func WriteFile(file string, data string) error {
	if file == "" {
		return errors.New("Please provide a file")
	}
	err := os.WriteFile(file, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
