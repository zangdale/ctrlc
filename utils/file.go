package utils

import (
	"os"
	"path/filepath"
)

func HomeDir() (string, error) {
	s, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(s, ".ctrlc")
	info, err := os.Stat(dir)
	if err == nil {
		if info.IsDir() {
			return dir, nil
		}
	}

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}
	return dir, nil
}
