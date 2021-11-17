package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{10, 10, 10, 10, 2}, 42},
		test{[]int{12, 24, 3}, 39},
		test{[]int{12, 13}, 25},
		test{[]int{22, 33}, 55},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
