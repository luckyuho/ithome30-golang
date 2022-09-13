package main

import (
	"fmt"
)

// Go的常數宣告與判斷式

const (
	Pi         = 3.14
	Gravity    = 9.8
	LightSpeed = 300000000
)

const (
	Black  = iota // 0
	Red           // 1
	Yellow        // 2
	Green         // 3
)

func main() {
	fmt.Println(Pi, Gravity, LightSpeed)
	fmt.Println(MultipliedBy8(3))
	fmt.Println(DividedBy8(16))

	day02Switch(Red)
	day02If(Green)
}

func day02If(n int) {
	if n == 1 {
		fmt.Println("Traffic lights are now red")
	} else if n == 2 {
		fmt.Println("Traffic lights are now yellow")
	} else if n == 3 {
		fmt.Println("Traffic lights are now green")
	} else {
		fmt.Println("Traffic lights are now black")
	}
}

func day02Switch(n int) {
	switch n {
	case Red:
		fmt.Println("Traffic lights are now red")
	case Yellow:
		fmt.Println("Traffic lights are now yellow")
	case Green:
		fmt.Println("Traffic lights are now green")
	default:
		fmt.Println("Traffic lights are now black")
	}
}

func MultipliedBy8(num int) int {
	return num << 3
}

func DividedBy8(num int) int {
	return num >> 3
}
