package bubble_sort

import (
	"fmt"
	"reflect"
	"testing"
)

func BubbleSort(numbers []int) []int {
	unsortedUntilIndex := len(numbers) - 1
	swap := true

	for swap {
		swap = false
		for i := 0; i < unsortedUntilIndex; i++ {
			first := numbers[i]
			second := numbers[i+1]

			if first > second {
				numbers[i] = second
				numbers[i+1] = first
				swap = true
			}
		}
		unsortedUntilIndex--
	}

	return numbers
}

func Equals(t *testing.T, module string, given string, should string, result interface{}, expected interface{}) {
	message := createTestMessage(module, given, should, result, expected)

	if reflect.DeepEqual(result, expected) {
		// fmt.Print(message)
	} else {
		t.Errorf(message)
	}
}

func createTestMessage(module string, given string, should string, result interface{}, expected interface{}) string {
	return fmt.Sprintf(
		`
%v
given:    %v
should:   %v
result:   %v
expected: %v
	`, module, given, should, result, expected)
}

// 原始測試寫法
// func TestBubbleSort(t *testing.T) {
// 	result := BubbleSort([]int{4, 2, 7, 1, 3})
// 	expected := []int{1, 2, 3, 4, 7}
// 	Equals(t, "BubbleSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, expected)

// 	result = BubbleSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
// 	expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	Equals(t, "BubbleSort()", "Given slice of integers", "Should return slice sorted in ascending order", result, expected)
// }
