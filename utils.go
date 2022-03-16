package zlog

import (
	"errors"
	"os"
)

// Check path
func pathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("A file with the same name exists")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
