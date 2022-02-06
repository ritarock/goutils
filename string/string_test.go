package string

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("Word", 3)
	want := "WordWordWord"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNumberToString(t *testing.T) {
	got := NumberToString(5)
	want := "5"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCapitalize(t *testing.T) {
	got := Capitalize("this is a pen.")
	want := "This is a pen."

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
