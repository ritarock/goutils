package file

import (
	"testing"
)

func TestRead(t *testing.T) {
	got := Read("./read.txt")
	want := "read text\n"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReadLine(t *testing.T) {
	got := ReadLine("./read.txt")
	want := "read text"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWrite(t *testing.T) {
	gotErr := Write("./write.txt", "write text")
	wantErr := false

	if (gotErr != nil) != wantErr {
		t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
	}
}
