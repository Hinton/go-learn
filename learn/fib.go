package learn

func Fib(i uint) uint {

	if i < 2 {
		return i
	}

	return Fib(i - 1) + Fib(i - 2)
}
