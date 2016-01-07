package learn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestFib(t *testing.T) {

	var tests = []uint{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}

	for i, num := range tests {
		assert.Equal(t, Fib(uint(i+1)), num)
	}

}
