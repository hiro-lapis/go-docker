package main

import (
	"flag"
	"fmt"
	"strconv"
)

var input = flag.Int("input", 0, "it is used to Fizzbuzz output")

func main() {
	flag.Parse()
	fmt.Println(*input)
	Fizzbuzz(*input)
}

func Fizzbuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(i)
}
