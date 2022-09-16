package main

import "fmt"

func main() {
	list1 := map[string]interface{}{
		"水果": "頻果",
		"數量": 5,
	}
	for i, v := range list1 {
		fmt.Printf("key: %v, value: %v \n", i, v)
	}
	list2 := map[interface{}]interface{}{
		"水果": "頻果",
		"數量": 5,
		7:    "保存期限",
	}
	for i, v := range list2 {
		fmt.Printf("key: %v, value: %v \n", i, v)
	}
}
