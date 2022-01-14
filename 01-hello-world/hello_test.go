package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got: %q | Want: %q", got, want)
		}
	}

	t.Run("Saying hello to demons", func(t *testing.T) {
		got := Hello("Inosuke")
		want := "Hello, Inosuke"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
