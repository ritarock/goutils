package file

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	path := "../../test_data/read_file.txt"
	result := Read(path)
	expect := "read file\n"

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestExists(t *testing.T) {
	path := "../../test_data/read_file.txt"
	result := Exists(path)
	expect := true

	if result != expect {
		t.Errorf("result = %v, want = %v", result, expect)
	}
}

func TestWrite(t *testing.T) {
	path := "../../test_data/write_file.txt"
	expect := "write file"
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

func TestRemoveFile(t *testing.T) {
	path := "../../test_data/write_file.txt"
	Write(path, "test")
	RemoveFile(path)
	result := Exists(path)
	expect := false

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
	oldPath := "../../test_data/old_path.txt"
	data := "tmp"
	newPath := "../../test_data/new_path.txt"
	Write(oldPath, data)
	Rename(oldPath, newPath)
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

func TestCopy(t *testing.T) {
	srcPath := "../../test_data/read_file.txt"
	destPath := "../../test_data/dest_file.txt"
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
