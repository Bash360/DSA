package fibonacci

import "testing"

func TestFibonnaci(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Fibonnaci Test 1 ", 1, 0},
		{"Fibonnaci Test 2 ", 20, 6765},
		{"Fibonnaci Test 3 ", 6, 8},
	}

	for _, v := range tests {
		result := Fibonacci(v.input)

		if result != v.want {
			t.Errorf("Name %s: want %d got %d", v.name, v.want, result)
		}
	}
}
