package helpers

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func GetFixturePath(name string) string {
	_, filename, _, _ := runtime.Caller(0)

	fixturePath := path.Join(path.Dir(filename), "..", "fixtures", name)

	_, err := os.Stat(fixturePath)

	if err != nil {
		panic(err)
	}

	return fixturePath
}

func ReadFixtureBytes(name string) []byte {
	path := GetFixturePath(name)
	contents, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return contents
}

func ReadFixtureString(name string) string {
	return string(ReadFixtureBytes(name))
}
