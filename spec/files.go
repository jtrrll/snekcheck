package e2e

import (
	"math/rand"
	"os"
	"path/filepath"
)

type File struct {
	Name     string
	Children []File
}

func CreateFile(path string, file File) string {
	path = filepath.Join(path, file.Name)

	if len(file.Children) > 0 {
		if err := os.Mkdir(path, os.ModeDir); err != nil {
			panic(err)
		}
		for _, child := range file.Children {
			CreateFile(path, child)
		}
	} else {
		if _, err := os.Create(path); err != nil {
			panic(err)
		}
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

const validChars string = "abcdefghijklmnopqrstuvwxyz0123456789_"

func ValidChars(length uint) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = validChars[rand.Intn(len(validChars))]
	}

	return string(buf)
}
