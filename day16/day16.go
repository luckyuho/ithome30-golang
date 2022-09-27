package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	val  int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.val++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.val
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		for {
			a.increment()
		}
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
