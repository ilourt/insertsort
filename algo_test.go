package binarysearch

import "testing"

var tests = []struct {
	list   []int
	value  int
	result bool
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, true},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 2, true},
	{[]int{1, 5, 7, 9, 12, 58}, 14, false},
}

func TestSearchSuccess(t *testing.T) {
	for _, test := range tests {
		res := Search(test.list, test.value)
		if res != test.result {
			t.Errorf("Error for list %v and value %d; result expected %t", test.list, test.value, test.result)
		}
	}
}
