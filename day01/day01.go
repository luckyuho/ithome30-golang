package main

import "fmt"

// Go的變數宣告與資料型態

var num int // num = 0
var (
	global_num int    // global_num = 0
	global_str string // global_str = ""
)

func main() {
	fmt.Println(num, global_num, global_str)
	num := 2022                        // num = 2022
	local_num, local_str := 2, "world" // local_num = 2, local_str = "world"
	fmt.Println(num, local_num, local_str)
}
