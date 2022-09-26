package main

import (
	"fmt"
	"sync"
)

// // version 3 資料蒐集完就 print
// func doWorker3(id int, c chan int, done chan bool) {
// 	for n := range c {
// 		fmt.Printf("Worker %d received %c \n", id, n)
// 		done <- true
// 	}
// }

// type worker3 struct {
// 	in   chan int
// 	done chan bool
// }

// func createWorker3(id int) worker3 {
// 	w := worker3{
// 		in:   make(chan int),
// 		done: make(chan bool),
// 	}
// 	go doWorker3(id, w.in, w.done)
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

// 	for i := 0; i < 10; i++ {
// 		<-workers[i].done // receive
// 	}

// 	for i, worker := range workers {
// 		worker.in <- 'A' + i
// 	}

// 	for i := 0; i < 10; i++ {
// 		<-workers[i].done // receive
// 	}
// }

// // version 4 資料蒐集完後統一 print
// func doWorker4(id int, c chan int, wg *sync.WaitGroup) {
// 	for n := range c {
// 		fmt.Printf("Worker %d received %c \n", id, n)
// 		// done <- true // 刪除
// 		wg.Done()
// 	}
// }

// type worker4 struct {
// 	in chan int
// 	wg *sync.WaitGroup
// }

// func createWorker4(id int, wg *sync.WaitGroup) worker4 {
// 	w := worker4{
// 		in: make(chan int),
// 		wg: wg,
// 	}
// 	go doWorker4(id, w.in, wg)
// 	return w
// }

// func chanDemo4() {
// 	var wg sync.WaitGroup

// 	var workers [10]worker4
// 	for i := 0; i < 10; i++ {
// 		workers[i] = createWorker4(i, &wg)
// 	}

// 	wg.Add(20)

// 	for i, worker := range workers {
// 		worker.in <- 'a' + i // send
// 	}

// 	// for i := 0; i < 10; i++ { // 刪除
// 	// 	<-workers[i].done // receive // 刪除
// 	// } // 刪除

// 	for i, worker := range workers {
// 		worker.in <- 'A' + i
// 	}

// 	// for i := 0; i < 10; i++ { // 刪除
// 	// 	<-workers[i].done // receive // 刪除
// 	// } // 刪除

// 	wg.Wait()

// }

// version 5 資料蒐集完後統一 print
func doWorker5(id int, w worker5) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c \n", id, n)
		// done <- true // 刪除
		w.done()
	}
}

type worker5 struct {
	in   chan int
	done func()
}

func createWorker5(id int, wg *sync.WaitGroup) worker5 {
	w := worker5{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker5(id, w)
	return w
}

func chanDemo5() {
	var wg sync.WaitGroup

	var workers [10]worker5
	for i := 0; i < 10; i++ {
		workers[i] = createWorker5(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i // send
	}

	// for i := 0; i < 10; i++ { // 刪除
	// 	<-workers[i].done // receive // 刪除
	// } // 刪除

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// for i := 0; i < 10; i++ { // 刪除
	// 	<-workers[i].done // receive // 刪除
	// } // 刪除

	wg.Wait()

}

func main() {
	// chanDemo3()
	// chanDemo4()
	chanDemo5()

}
