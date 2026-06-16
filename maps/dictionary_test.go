package main

import "testing"

func TestSearch(t *testing.T) {
	dicionary := map[string]string{"test": "this is just a test"}

	got := Search(dicionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
