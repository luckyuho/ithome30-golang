package main

import (
	"fmt"
)

// // version 1 需要 TimeSleep 的寫法，則會 print 的結果不完整
// func worker1(id int, c chan int) {
// 	for n := range c {
// 		fmt.Printf("Worker %d received %c \n", id, n)
// 	}
// }

// func createWorker1(id int) chan<- int {
// 	c := make(chan int)
// 	go worker1(id, c)
// 	return c
// }

// func chanDemo1() {
// 	var channels [10]chan<- int
// 	for i := 0; i < 10; i++ {
// 		channels[i] = createWorker1(i)
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'a' + i
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'A' + i
// 	}

// 	// time.Sleep(time.Millisecond)
// }

// // version 2 不需要 TimeSleep
// // 但會按照順序打印，表示是單純按照順序執行，沒有發揮 goroutine 的功能
// func doWorker2(id int, c chan int, done chan bool) {
// 	for n := range c {
// 		fmt.Printf("Worker %d received %c \n", id, n)
// 		done <- true
// 	}
// }

// type worker2 struct {
// 	in   chan int
// 	done chan bool
// }

// func createWorker2(id int) worker2 {
// 	w := worker2{
// 		in:   make(chan int),
// 		done: make(chan bool),
// 	}
// 	go doWorker2(id, w.in, w.done)
// 	return w
// }

// func chanDemo2() {
// 	var workers [10]worker2
// 	for i := 0; i < 10; i++ {
// 		workers[i] = createWorker2(i)
// 	}

// 	for i := 0; i < 10; i++ {
// 		workers[i].in <- 'a' + i // send
// 		<-workers[i].done        // receive，主要是因為被這行塞住了，所以導致會按順序執行
// 	}

// 	for i := 0; i < 10; i++ {
// 		workers[i].in <- 'A' + i
// 		<-workers[i].done
// 	}
// }

// version 3 資料蒐集完就 print
func doWorker3(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		done <- true
	}
}

type worker3 struct {
	in   chan int
	done chan bool
}

func createWorker3(id int) worker3 {
	w := worker3{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker3(id, w.in, w.done)
	return w
}

func chanDemo3() {
	var workers [10]worker3
	for i := 0; i < 10; i++ {
		workers[i] = createWorker3(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i // send
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done // receive
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done // receive
	}
}

func main() {
	// chanDemo1()
	// chanDemo2()
	chanDemo3()
}
