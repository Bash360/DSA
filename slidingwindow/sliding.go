package slidingwindow

func Sliding(arr []int, windowSize int) int {

	window := arr[:windowSize]
	windowSum := 0
	for _, v := range window {
		windowSum += v
	}
	max := windowSum
	right := windowSize
	left := 0
	for right < len(arr) {

		windowSum = (windowSum - arr[left]) + arr[right]
		if windowSum > max {
			max = windowSum
		}
		right++
		left++
	}

	return max
}

// 0(n)

func Sum(arr []int, windowSize int) []int {
	window := arr[:windowSize]
	var sum int
	result := make([]int, 0)

	for _, v := range window {
		sum += v
	}

	result = append(result, sum)

	left := 0
	right := windowSize

	for right < len(arr) {

		sum = sum - arr[left] + arr[right]

		right++
		left++
		result = append(result, sum)

	}

	return result

}

// 0(n)



