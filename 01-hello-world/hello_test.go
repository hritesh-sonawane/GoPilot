package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got: %q | Want: %q", got, want)
		}
	}

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("to a demon", func(t *testing.T) {
		got := Hello("Akaza", "")
		want := "Hello, Akaza"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Inosuke", spanish)
		want := "Hola, Inosuke"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Zenitsu", french)
		want := "Bonjour, Zenitsu"
		assertCorrectMessage(t, got, want)
	})
}
