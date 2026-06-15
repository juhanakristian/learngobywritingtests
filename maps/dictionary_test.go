package main

import "testing"

func TestSearch(t *testing.T) {
	dicionary := map[string]string{"test": "this is just a test"}

	got := Search(dicionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
