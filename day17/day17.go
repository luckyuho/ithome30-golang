package main

import (
	"fmt"
	"time"
)

// // version 3 資料蒐集完就 print
// func doWorker3(id int, c chan int) {
// 	for n := range c {
// 		fmt.Printf("Worker %d received %c \n", id, n)
// 	}
// 	close(c)
// }

// type worker3 struct {
// 	in chan int
// 	// done chan bool
// }

// func createWorker3(id int) worker3 {
// 	w := worker3{
// 		in: make(chan int, 2),
// 		// done: make(chan bool),
// 	}
// 	go doWorker3(id, w.in)
// 	return w
// }

// func chanDemo3() {
// 	var workers [10]worker3
// 	for i := 0; i < 10; i++ {
// 		workers[i] = createWorker3(i)
// 	}

// 	for i, worker := range workers {
// 		worker.in <- 'a' + i // send
// 	}

// 	for i, worker := range workers {
// 		worker.in <- 'A' + i
// 	}
// }

// func main() {
// 	chanDemo3()
// }

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	tm := time.After(3 * time.Second)

	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		case <-tm:
			return
		default:
			fmt.Println("No value received")
		}
	}
}
