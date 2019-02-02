package lib

import (
	"io/ioutil"
	"os"
)

// Read reads and file and return its content
func Read(fileName string) (string, error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	r, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}
	return string(r), nil
}
