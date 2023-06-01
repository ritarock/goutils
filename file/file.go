package file

import (
	"bufio"
	"io"
	"os"
)

func Read(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadLine(path string) (string, error) {
	var line string
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		line = s.Text()
	}
	if s.Err() != nil {
		return "", err
	}

	return line, nil
}

func Write(path, data string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}
