package main

import (
	"fmt"
	"strconv"
)

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

	res4 := fizzbuzz(4)
	fmt.Println(res4)

	res5 := fizzbuzz(5)
	fmt.Println(res5)
}

func fizzbuzz(num int) string {
	res := ""

	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			res = res + "fizz"
		} else if i%5 == 0 {
			res = res + "buzz"
		} else {
			res = res + strconv.Itoa(i)
		}
		res = res + " "
	}
	return res
}
