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
			path:     "../testdata/test.tx",
			hasError: true,
		},
	}

	for _, test := range tests {
		got, err := Read(test.path)
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		}
	}
}

func TestWrite(t *testing.T) {
	tmp, err := os.CreateTemp("./", "testfile")
	if err != nil {
		t.Fatalf("failed TestWrite")
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	tests := []struct {
		path string
		data string
	}{
		{
			path: tmp.Name(),
			data: "",
		},
	}

	for _, test := range tests {
		err := Write(test.path, test.data)
		assert.NoError(t, err)
	}
}
