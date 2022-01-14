package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Zenitsu")
	want := "Hello, Zenitsu"

	if got != want {
		t.Errorf("Got: %q | Want: %q", got, want)
	}
}
