package file

import (
	"testing"
)

func assertNoError(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, gotErr error, wantErr bool) {
	t.Helper()
	if (gotErr != nil) != wantErr {
		t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
	}
}

func TestRead(t *testing.T) {
	t.Run("exist read text", func(t *testing.T) {
		got, _ := Read("./read.txt")
		want := "read text\n"
		assertNoError(t, got, want)
	})

	t.Run("NOT exist read text", func(t *testing.T) {
		_, gotErr := Read("./read.tx")
		wantErr := true
		assertError(t, gotErr, wantErr)
	})
}

func TestReadLine(t *testing.T) {
	t.Run("exist read text", func(t *testing.T) {
		got, _ := ReadLine("./read.txt")
		want := "read text"
		assertNoError(t, got, want)
	})

	t.Run("NOT exist read text", func(t *testing.T) {
		_, gotErr := ReadLine("./read.tx")
		wantErr := true
		assertError(t, gotErr, wantErr)
	})
}

func TestWrite(t *testing.T) {
	gotErr := Write("./write.txt", "write text")
	wantErr := false

	if (gotErr != nil) != wantErr {
		t.Errorf("gotErr %v, wantErr %v", gotErr, wantErr)
	}
}
