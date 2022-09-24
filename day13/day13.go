package main

import (
	"fmt"
	"time"
)

// 情況 1
func chanDemo1() {
	// chan type 是 channel 的資料型態
	// 可以是 chan int 也可以是 chan string 等等
	c := make(chan int, 2)

	c <- 1 // 發資料到 channel
	c <- 2
	n := <-c // 收 channel 資料

	fmt.Println(n)
	n = <-c
	fmt.Println(n)
}

// 情況 2
func chanDemo2() {
	// chan type 是 channel 的資料型態
	// 可以是 chan int 也可以是 chan string 等等
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
	c <- 7
	c <- 8
	c <- 9
	c <- 0
	// close(c)
}

func chanDemo3() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = makeChan3()
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
}

func makeChan3() chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	return c
}

func chanDemo4() {
	// 此為 channel 純收資料
	// 如果想要變成 channel 純發資料，則寫成 <-chan int
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = makeChan4()
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}
}

func makeChan4() chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	return c
}

func main() {
	// chanDemo1()
	chanDemo2()
	// chanDemo3()
	// chanDemo4()

	// 多下面這行等待時間是要避免 goroutine 間的資料還沒傳遞完
	// 程式碼就已經結束運行
	time.Sleep(time.Millisecond)
}
