package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func Write(path, data string) error {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write([]byte(data))
	return err
}

func RemoveFile(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

func MakeDir(path string) {
	err := os.Mkdir(path, 0777)
	if err != nil {
		panic(err)
	}
}

func RemoveDir(path string) {
	RemoveFile(path)
}

func ReadLine(path string) string {
	var line string
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)

	for s.Scan() {
		line = s.Text()
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return line
}
