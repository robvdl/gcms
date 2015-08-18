package util

import "os"

// Exists checks whether a file exists or not
func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
