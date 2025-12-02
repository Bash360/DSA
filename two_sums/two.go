package twosums

func TwoSums(arr []int, target int) []int {
	numsFrequency := make(map[int]int, len(arr))
	var result []int

	for _, v := range arr {
		numsFrequency[v] = v
	}

	for _, v := range arr {
		possibleValue := target - v
		_, ok := numsFrequency[possibleValue]
		if ok {
			result = []int{v, possibleValue}

		}
	}
	return result
}

// 0(n)
