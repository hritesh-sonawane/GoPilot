package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Giyu")

	got := buffer.String()
	want := "Hello, Giyu"

	if got != want {
		t.Errorf("Got %q | Want %q", got, want)
	}
}
