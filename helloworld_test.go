package main

import "testing"

func TestHelloAndrew(test *testing.T) {
	got := Hello("Andrew!")
	want := "Hello, Andrew!"

	if got != want {
		test.Errorf("got %q want %q", got, want)
	}
}