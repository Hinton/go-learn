package learn

import (
	"strconv"
)

const (
	fizz     = "fizz"
	buzz     = "buzz"
	fizzBuzz = "fizzbuzz"
)

func FizzBuzz(i int) string {

	if i%3 == 0 && i%5 == 0 {
		return fizzBuzz
	}

	if i%3 == 0 {
		return fizz
	}

	if i%5 == 0 {
		return buzz
	}

	return strconv.Itoa(i)

}
