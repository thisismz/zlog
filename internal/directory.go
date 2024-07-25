package internal

import (
	"errors"
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("path exists but is not a directory")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			fmt.Printf("create directory: %s\n", v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				fmt.Printf("create directory error: %v\n", err)
				return err
			}
		}
	}
	return err
}
