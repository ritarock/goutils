package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     string
		hasError bool
	}{
		{
			name:     "exists read file",
			path:     "../testdata/test.txt",
			want:     "This is test file.\n",
			hasError: false,
		},
		{
			name:     "not exists read file",
			path:     "../testdata/readfile.tx",
			hasError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Read(test.path)
			if test.hasError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, test.want, got)
				assert.NoError(t, err)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	tmpfile, err := os.CreateTemp("./", "testfile")
	if err != nil {
		t.Fatalf("failed TestWrite")
	}
	defer os.Remove(tmpfile.Name())
	defer tmpfile.Close()

	tests := []struct {
		name string
		path string
		data string
		want string
	}{
		{
			name: "exists read file",
			path: tmpfile.Name(),
			data: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Write(test.path, test.data)
			assert.NoError(t, err)
		})
	}
}
