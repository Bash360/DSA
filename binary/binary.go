package binary

func BinarySearch(arr []int, search int) int {

	left := 0
	right := len(arr) - 1
	var mid int

	for left <= right {
		mid = left + right/2

		midValue := arr[mid]

		switch {
		case search > midValue:
			left = mid + 1
		case search < midValue:
			right = mid - 1
		default:
			return mid
		}

	}
	return -1
}

// 0(logn)
