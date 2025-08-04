package e2e

import (
	"math/rand"
	"os"
	"path/filepath"
)

func CreateFile(path string, name string) string {
	path = filepath.Join(path, name)
	if _, err := os.Create(path); err != nil {
		panic(err)
	}

	return path
}

func CreateDirectory(path string, name string) string {
	path = filepath.Join(path, name)
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		panic(err)
	}

	return path
}

var TestDir string = filepath.Join(os.TempDir(), "snekcheck_e2e")

func ResetTestDir() {
	if err := os.RemoveAll(TestDir); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(TestDir, os.ModePerm); err != nil {
		panic(err)
	}
}

const validChars string = "abcdefghijklmnopqrstuvwxyz0123456789"

func ValidChars(length uint) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = validChars[rand.Intn(len(validChars))]
	}

	return string(buf)
}
