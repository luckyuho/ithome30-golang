package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

// type aa = func(int) int
// type aab func(int) int

// func test(base int) aa {
// 	return func(int) int {
// 		return 1
// 	}
// }

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	// a1 := adder()
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("0 + 1 + ... + %d = %d \n", i, a(i))
	// }
	// fmt.Println(a1(5), &a1)
	// fmt.Println(a1(10), &a1)
	// a1 = adder()
	// fmt.Println(a1(5), &a1)
	// fmt.Println(a1(10), &a1)

	a2 := adder2(0)
	result, a2 := a2(5)
	fmt.Println(result)
	result, a2 = a2(10)
	fmt.Println(result)
}

// 自己練習
type iAdder_2 func(int) (int, iAdder_2)

func adder_2(base int) iAdder_2 {
	return func(v int) (int, iAdder_2) {
		return v + base, adder_2(v + base)
	}
}
