package main

import (
	"testing"
)

/*
	1. fizzbuzz(1)  should output 1 DONE
	2. fizzbuzz(3)  should output 1 2 fizz DONE
	3. fizzbuzz(5) should output 1 2 fizz 4 buzz
	4. fizzbuzz(15) should output 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz
*/

func TestFizzbuzz(t *testing.T) {

	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("number 1 passed to the function should return the number", func(t *testing.T) {
		got := fizzbuzz(1)
		want := "1"

		if got != want {
			assertMessage(t, got, want)
		}
	})

	t.Run("3 modulus of 3 should return fizz", func(t *testing.T) {
		got := fizzbuzz(3)
		want := "1 2 fizz"

		if got != want {
			assertMessage(t, got, want)
		}
	})

	t.Run("5 modulus of 5 should return buzz", func(t *testing.T) {
		got := fizzbuzz(5)
		want := "1 2 fizz 4 buzz"

		if got != want {
			assertMessage(t, got, want)
		}
	})

}
