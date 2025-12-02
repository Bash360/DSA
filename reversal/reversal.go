package reversal

func Reversal(word string) string {
	left := 0
	right := len(word) - 1
	result := make([]byte, len(word))

	for left <= right {
       
		result[left] = word[right]
		result[right] = word[left]

		left++
		right--
	}
	return string(result)
}
