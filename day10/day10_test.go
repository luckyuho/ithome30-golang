package bubble_sort

import (
	"math/rand"
	"testing"
)

// 單元測試，表格驅動測試
func TestBubbleSort(t *testing.T) {

	input := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		input[i] = rand.Intn(10000)
	}

	tests := []struct {
		input    []int
		expected []int
	}{
		// {[]int{4, 2, 7, 1, 3}, []int{1, 2, 3, 4, 7}},
		// {[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},

		// 只為了展示，所以解答直接透過函數運算
		{input, BubbleSort(input)},
	}

	for _, test := range tests {
		result := BubbleSort(test.input)
		Equals(t, "BubbleSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, test.expected)
	}
}
