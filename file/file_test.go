package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	tmpFile, err := os.CreateTemp("./", "test_read_*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	content := "This is test file"
	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)

	err = tmpFile.Close()
	assert.NoError(t, err)

	tests := []struct {
		name     string
		path     string
		want     string
		hasError bool
	}{
		{
			name:     "succeed",
			path:     tmpFile.Name(),
			want:     "This is test file",
			hasError: false,
		},
		{
			name:     "failed: not exists read file",
			path:     "tmpFile",
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
	tmpFile, err := os.CreateTemp("./", "test_read_*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	err = tmpFile.Close()
	assert.NoError(t, err)

	tests := []struct {
		name string
		path string
		data string
	}{
		{
			name: "succeed",
			path: tmpFile.Name(),
			data: "",
		},
	}

	for _, test := range tests {
		err := Write(test.path, test.data)
		assert.NoError(t, err)
	}
}
