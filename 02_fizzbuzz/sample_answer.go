package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(number int) (string) {
	if number % 15 == 0 {
	  return "FizzBuzz"
	}
	if number % 5 == 0 {
	  return "Buzz"
	}
	if number % 3 == 0 {
	   return "Fizz"
	}
	return strconv.Itoa(number)
}

func main() {
	fmt.Println(fizzbuzz(1))
	fmt.Println(fizzbuzz(2))
	fmt.Println(fizzbuzz(3))
	fmt.Println(fizzbuzz(4))
	fmt.Println(fizzbuzz(5))
	fmt.Println(fizzbuzz(6))
	fmt.Println(fizzbuzz(7))
	fmt.Println(fizzbuzz(8))
	fmt.Println(fizzbuzz(9))
	fmt.Println(fizzbuzz(10))
	fmt.Println(fizzbuzz(11))
	fmt.Println(fizzbuzz(12))
	fmt.Println(fizzbuzz(13))
	fmt.Println(fizzbuzz(14))
	fmt.Println(fizzbuzz(15))
}
