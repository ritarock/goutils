package file

import (
	"os"
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	path := "../../test_data/test_read.txt"
	result := Read(path)
	expect := "read text"
	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestExists(t *testing.T) {
	path := "../../test_data/test.txt"
	result := Exists(path)
	expect := true
	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestWrite(t *testing.T) {
	path := "../../test_data/test_write.txt"
	expect := "write text"
	Write(path, expect)
	result := Read(path)

	defer func() {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatal("err = \v", err)
		}
	}()

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestMakeDir(t *testing.T) {
	path := "../../test_data/tmp"
	MakeDir(path)
	result := Exists(path)
	expect := true

	defer func() {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatal("err = \v", err)
		}
	}()

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestRemoveFile(t *testing.T) {
	path := "../../test_data/test_write.txt"
	data := "text"
	Write(path, data)
	RemoveFile(path)
	result := Exists(path)
	expect := false
	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestRemoveDir(t *testing.T) {
	path := "../../test_data/tmp"
	MakeDir(path)
	RemoveDir(path)
	result := Exists(path)
	expect := false
	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestRename(t *testing.T) {
	path := "../../test_data/test_write.txt"
	data := "text"
	newPath := "../../test_data/new_path.txt"
	Write(path, data)
	Rename(path, newPath)
	result := Exists(newPath)
	expect := true

	defer func() {
		err := os.RemoveAll(newPath)
		if err != nil {
			t.Fatal("err = \v", err)
		}
	}()

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestReadDir(t *testing.T) {
	path := "../../test_data"
	result := ReadDir(path)
	expect := []string{"test_read.txt"}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("result = %v, want %v", result, expect)
	}
}

func TestCopy(t *testing.T) {
	srcPath := "../../test_data/test_read.txt"
	destPath := "../../test_data/copy.txt"
	Copy(srcPath, destPath)

	result := Exists(destPath)
	expect := true

	defer func() {
		err := os.RemoveAll(destPath)
		if err != nil {
			t.Fatal("err = \v", err)
		}
	}()

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}
