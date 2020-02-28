package utils

import (
	"os"
)

// IsFile check file exist
func IsFile(name string) (bool, error) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// IsDir check is Dir
func IsDir(dir string) (bool, error) {
	d, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if !os.IsNotExist(err) {
			return false, err
		}
	}
	return d.IsDir(), nil
}
