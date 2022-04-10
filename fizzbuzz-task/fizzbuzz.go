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
		if i%3 == 0 && i%5 == 0 {
			fizzbuzzResult = fizzbuzzResult + "fizzbuzz"
		} else if i%5 == 0 {
			fizzbuzzResult = fizzbuzzResult + "buzz"
		} else if i%3 == 0 {
			fizzbuzzResult = fizzbuzzResult + "fizz"
		} else {
			fizzbuzzResult = fizzbuzzResult + strconv.Itoa(i)
		}
		if i != number {
			fizzbuzzResult = fizzbuzzResult + " "
		}
	}
	return fizzbuzzResult
}
