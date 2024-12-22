package zlog

import (
	"errors"
	"os"
)

func pathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("not a directory")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
