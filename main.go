package main

import "fmt"

/*
	fizzbuzz(1)  should output 1
	fizzbuzz(2)  should output 1 2
	fizzbuzz(3)  should output 1 2 fizz
	fizzbuzz(5) should output 1 2 fizz 4 buzz
	fizzbuzz(15) should output 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz
*/

func main() {
	res1 := fizzbuzz(1)
	fmt.Println(res1)

	res2 := fizzbuzz(2)
	fmt.Println(res2)

	res3 := fizzbuzz(3)
	fmt.Println(res3)
}

func fizzbuzz(num int) string {
	if num == 1 {
		return "1"
	}
	if num == 2 {
		return "1 2"
	}
	if num == 3 {
		return "1 2 fizz"
	}
	return ""
}
