package string

import (
	"testing"
)

func assert(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("Hello", 3)
	want := "HelloHelloHello"

	assert(t, got, want)
}

func TestNumberToString(t *testing.T) {
	got := NumberToString(5)
	want := "5"

	assert(t, got, want)
}

func TestCapitalize(t *testing.T) {
	got := Capitalize("hello world")
	want := "Hello world"

	assert(t, got, want)
}
