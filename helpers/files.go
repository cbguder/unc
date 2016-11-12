package helpers

import "io/ioutil"

func ReadFileBytes(path string) []byte {
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return contents
}

func ReadFileString(path string) string {
	return string(ReadFileBytes(path))
}
