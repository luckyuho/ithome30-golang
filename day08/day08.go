package main

import (
	"fmt"
	"time"
)

// 情境1: 簡易的 error handling
// STEP 1：定義客製化的 error struct
type MyError struct {
	Time    time.Time
	Message string
}

type YourError struct {
	Time    time.Time
	Message string
}

// STEP 2：定義能夠屬於 Error Interface 的方法
func (e *MyError) Error() string {
	return fmt.Sprintf("喔耶，已經%v了，%s", e.Time.Format("15:04:05"), e.Message)
}

func (e *YourError) Error() string {
	return fmt.Sprintf("已經%v了，%s", e.Time.Format("15:04:05"), e.Message)
}

// STEP 3：拋出錯誤的函式
func errHandler(status bool) (error, bool) {
	if status {
		return nil, true
	}
	return &MyError{
		Time:    time.Now(),
		Message: "我要來耍廢了",
	}, false
}

func yourErrHandler(status bool) (error, bool) {
	if status {
		return nil, true
	}
	return &YourError{
		Time:    time.Now(),
		Message: "換你耍廢了",
	}, false
}

// STEP 4：使用 fmt.Println 即可取得錯誤拋出的訊息
func main() {
	// result := run(errHandler)
	// result(true)
	// result(false)
	isMyError(errHandler)
	isMyError(yourErrHandler)
}

// PS 只是讓程式碼比較好看的改寫，不一定需要
type errType func(bool) (error, bool)

// PS 只是讓程式碼比較好看的改寫，不一定需要
func run(op errType) func(bool) {
	return func(status bool) {
		if err, ok := op(status); ok {
			fmt.Println("ok")
		} else {
			fmt.Println(err)
		}
	}
}

func isMyError(op errType) {
	err, _ := op(false)
	if err != nil {
		switch err.(type) {
		case *MyError:
			fmt.Println("MyError: ", err)
		case *YourError:
			fmt.Println("YourError: ", err)
		default:
			fmt.Println("沒有定義是誰的錯")
		}
	}
}
