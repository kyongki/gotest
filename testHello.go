package main

import "testing"

func testHello(t *testing.T) {
	got := Hello("James")
	want := "Hello, john"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
