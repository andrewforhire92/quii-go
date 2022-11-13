package main

import "testing"

func TestHelloAndrew(test *testing.T) {
	test.Run("saying hello to people", func(test *testing.T) {
		got := Hello("Andrew!")
		want := "Hello, Andrew!"

		// if got != want {
		// 	test.Errorf("got %q want %q", got, want)
		// }

		// ** We move the above code into its own function ** :

		assertCorrectMessage(test, got, want)
	})
	test.Run("say 'Hello, World' when an empty string is supplied", func(test *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertCorrectMessage(test, got, want)
	})
}

func assertCorrectMessage(test testing.TB, got, want string) {
	test.Helper()
	if got != want {
		test.Errorf("got %q want %q", got, want)
	}
}
