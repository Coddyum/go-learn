package main

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat x time a string entry", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}