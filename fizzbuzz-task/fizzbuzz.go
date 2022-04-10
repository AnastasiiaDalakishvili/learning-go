package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := fizzbuzz(1)
	fmt.Println(result)

	result3 := fizzbuzz(3)
	fmt.Println(result3)
}

func fizzbuzz(number int) string {
	fizzbuzzResult := ""

	for i := 1; i <= number; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fizzbuzzResult = fizzbuzzResult + "fizzbuzz"
		case i%3 == 0:
			fizzbuzzResult = fizzbuzzResult + "fizz"
		case i%5 == 0:
			fizzbuzzResult = fizzbuzzResult + "buzz"
		default:
			fizzbuzzResult = fizzbuzzResult + strconv.Itoa(i)
		}
		if i != number {
			fizzbuzzResult = fizzbuzzResult + " "
		}
	}
	return fizzbuzzResult
}
