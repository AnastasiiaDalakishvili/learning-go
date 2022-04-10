package main

import "testing"

/*
	fizzbuzz(1)  should output 1
	fizzbuzz(2)  should output 1 2
	fizzbuzz(3)  should output 1 2 fizz
	fizzbuzz(5) should output 1 2 fizz 4 buzz
	fizzbuzz(15) should output 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz
*/

func TestFizzbuzz(t *testing.T) {
	got := fizzbuzz(1)
	want := 1

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
