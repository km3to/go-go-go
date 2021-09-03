package softuni

var data = [1000]int{}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	if data[n] == 0 {
		data[n] = Fibonacci(n-1) + Fibonacci(n-2)
	}

	return data[n]
}
