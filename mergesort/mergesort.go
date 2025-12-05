package mergesort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](arr ...T) []T {

	if len(arr) <= 1 {

		return arr
	}

	mid := len(arr) / 2

	left := Sort(arr[:mid]...)
	right := Sort(arr[mid:]...)

	return merge(left, right)
}


func merge[T constraints.Ordered](left []T, right []T) []T {
	result := make([]T, 0)
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
