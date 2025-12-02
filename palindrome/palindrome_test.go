package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{{"Palindrome Test 1 ", "racecar", true}, {"Palindrome Test 2 ", "did", true}, {"Palindrome Test 3 ", "child", false}}

	for _, v := range tests {
		result := Palindrome(v.input)

		if result != v.want {
			t.Errorf("Name %s: want %t got %t", v.name, v.want, result)
		}

	}
}
