package util

import (
	"os"
)

func ReadFile(filename string) string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(contents)
}
