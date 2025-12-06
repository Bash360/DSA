package fibonacci

func Fibonacci(num int) int {
	left := 0
	right := 1
	var result int
	for range num - 1 {
		result = left + right
		left = right
		right = result

	}

	return result

}

func FibRecursion(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return FibRecursion(num-2) + FibRecursion(num-1)

}
