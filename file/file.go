package file

import (
	"bufio"
	"io"
	"log"
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
		log.Println(err)
	}
	return string(b)
}

func ReadLine(path string) string {
	var line string
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		line = s.Text()
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return line
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
