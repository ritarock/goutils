package file

import (
	"io/ioutil"
	"log"
	"os"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Write(path, data string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(([]byte)(data))
}

func RemoveFile(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}

func MakeDir(path string) {
	err := os.Mkdir(path, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveDir(path string) {
	RemoveFile(path)
}

func Rename(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func Copy(srcPath, destPath string) {
	err := os.Link(srcPath, destPath)
	if err != nil {
		log.Fatal(err)
	}
}
