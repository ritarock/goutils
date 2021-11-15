package string

import (
	"testing"
)

func TestCombine(t *testing.T) {
	got := Combine("Hello", " ", "World!")
	want := "Hello World!"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("word", 3)
	want := "wordwordword"

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
