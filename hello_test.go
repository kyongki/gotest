package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello with name", func(t *testing.T) {
		got := Hello("James")
		want := "Hello, James"
		assertCorrectMsg(t, got, want)
	})
	t.Run("hello with default", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, Everyone"
		assertCorrectMsg(t, got, want)
	})
	t.Run("support Spanish greeting", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, but want %q", got, want)
	}
}
