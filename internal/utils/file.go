package utils

import (
	"errors"
	"os"
)

func CreateFileIfNotExists(name string) error {
	_, err := os.Stat(name)
	if errors.Is(err, os.ErrNotExist) {
		os.WriteFile(name, []byte("[]"), 0666)
	}
	return err
}
