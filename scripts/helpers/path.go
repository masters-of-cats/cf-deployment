package helpers

import (
	"os"
	"path/filepath"
)

func SetPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Abs(filepath.Join(wd, "..", ".."))
}