package learn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoFizzBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(1), "1")
	assert.Equal(t, FizzBuzz(2), "2")
	assert.Equal(t, FizzBuzz(4), "4")
	assert.Equal(t, FizzBuzz(7), "7")
	assert.Equal(t, FizzBuzz(8), "8")
}

func TestFizz(t *testing.T) {
	const fizz = "fizz"
	assert.Equal(t, FizzBuzz(3), fizz)
	assert.Equal(t, FizzBuzz(6), fizz)
	assert.Equal(t, FizzBuzz(9), fizz)
	assert.Equal(t, FizzBuzz(12), fizz)
	assert.Equal(t, FizzBuzz(18), fizz)
}

func TestBuzz(t *testing.T) {
	const buzz = "buzz"
	assert.Equal(t, FizzBuzz(5), buzz)
	assert.Equal(t, FizzBuzz(10), buzz)
	assert.Equal(t, FizzBuzz(20), buzz)
	assert.Equal(t, FizzBuzz(25), buzz)
}

func TestFizzBuzz(t *testing.T) {
	const fizzBuzz = "fizzbuzz"
	assert.Equal(t, FizzBuzz(15), fizzBuzz)
}
