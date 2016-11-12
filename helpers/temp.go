package helpers

import (
	"os"
	"io/ioutil"
)

func CreateTempFile(prefix string) *os.File {
	file, err := ioutil.TempFile("", prefix)

	if err != nil {
		panic(err)
	}

	return file
}

func DiscardTempFile(file *os.File) {
	err := file.Close()

	if err != nil {
		panic(err)
	}

	err = os.Remove(file.Name())

	if err != nil {
		panic(err)
	}
}
