package fibonacci

func Fibonacci(num int) int {
	num1 := 0
	num2 := 1
	var result int
	for range num - 1 {
		result = num1 + num2
		num1 = num2
		num2 = result

	}

	return result

}

// 0(n)
