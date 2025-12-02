package binary

import (
	"testing"
)

func TestBinary(t *testing.T) {
	type input struct {
		arr    []int
		target int
	}
	tests := []struct {
		name  string
		input input
		want  int
	}{{"Binary Test 1", input{[]int{1, 2, 3, 4, 5}, 5}, 4},
		{"Binary Test 2", input{[]int{1, 4, 7, 9, 12, 18, 22}, 1}, 0},
		{"Binary Test 3", input{[]int{2, 5, 9, 14, 20, 31}, 14}, 3},
		{"Binary Test 4", input{[]int{3, 6, 8, 11, 15, 19}, 19}, 5},
	}

	for _, v := range tests {
		answer := BinarySearch(v.input.arr, v.input.target)
		if answer != v.want {
			t.Errorf("Name %s: want %d got %d", v.name, v.want, v.input.target)
		}
	}

}
