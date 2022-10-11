package main

import "fmt"

// 創建 2 維矩陣
func create2DimArray(m, n int) [][]int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)

		// 裡面塞值，來確認是哪一個 row 的資料
		for j := 0; j < n; j++ {
			if j == 0 {
				arr[i][j] = i
			} else {
				arr[i][j] = j
			}
		}
	}
	return arr
}

// 顯示 2 維陣列的內容物
func show2DimArray(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

// 對 1 個 row 裡面的每個值 +1
func add1Row(arr1Dim []int, c chan int) {
	for i, v := range arr1Dim {
		// 查看併發順序
		if i == 0 {
			fmt.Println("併發讀取的值", v)
		}
		c <- v + 1 // 每個值 +1
	}
}

// 對 1 個 row 裡面的每個值 +1，且輸出
func add1RowAndShow(arr [][]int, m, n int) {
	var c [10]chan int // 這邊必須是靜態的，必須給定 const 值，不能用 m 取代

	// 創建 2 維且長度為 10 的通道
	for i := range c {
		c[i] = make(chan int, n)
	}

	// 併發讀取與寫入
	for i := 0; i < m; i++ {
		go add1Row(arr[i], c[i])
	}

	// // 無法併發讀取，因為 c[i] 限制了他的順序
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Printf("%d ", <-c[i])
	// 	}
	// 	fmt.Println()
	// }

	// // 併發輸出，錯誤形式
	// for i := 0; i < m; i++ {
	// 	go showChanInfoBad(n, c[i])
	// }

	// 既然上面的方法會導致有順序性
	// 那我們讓 print 方法也變成併發的形式
	over := make(chan int, n)
	for i := 0; i < m; i++ {
		go showChanInfo(n, c[i], over)
	}
	for i := 0; i < m; i++ {
		<-over
	}
}

// 為了併發顯示的方式 (錯)
func showChanInfoBad(n int, c chan int) {
	for j := 0; j < n; j++ {
		fmt.Printf("%d ", <-c)
	}
	fmt.Println()
}

// 為了併發顯示的方式 (對)
func showChanInfo(n int, c, over chan int) {
	for j := 0; j < n; j++ {
		fmt.Printf("%d ", <-c)
	}
	fmt.Println()
	over <- 0
}

func main() {
	m, n := 10, 10
	arr := create2DimArray(m, n)

	showTitle("=== 顯示 2 維 arr 內容物 ===")
	show2DimArray(arr)

	showTitle("=== goroutine 對 2 維每筆資料 +1 與 輸出結果 ===")
	add1RowAndShow(arr, m, n)

}

func showTitle(s string) {
	fmt.Println(s)
}

// ~~~
// === 顯示 2 維 arr 內容物 ===
// 0 1 2 3 4 5 6 7 8 9
// 1 1 2 3 4 5 6 7 8 9
// 2 1 2 3 4 5 6 7 8 9
// 3 1 2 3 4 5 6 7 8 9
// 4 1 2 3 4 5 6 7 8 9
// 5 1 2 3 4 5 6 7 8 9
// 6 1 2 3 4 5 6 7 8 9
// 7 1 2 3 4 5 6 7 8 9
// 8 1 2 3 4 5 6 7 8 9
// 9 1 2 3 4 5 6 7 8 9
// === goroutine 對 2 維每筆資料 +1 與 輸出結果 ===
// 1 2 3 4 5 6 7 8 9 10
// 2 2 3 4 5 6 7 8 9 10
// 3 2 3 4 5 6 7 8 9 10
// 4 2 3 4 5 6 7 8 9 10
// 5 2 3 4 5 6 7 8 9 10
// 6 2 3 4 5 6 7 8 9 10
// 7 2 3 4 5 6 7 8 9 10
// 8 2 3 4 5 6 7 8 9 10
// 9 2 3 4 5 6 7 8 9 10
// 10 2 3 4 5 6 7 8 9 10
// === goroutine 對 2 維每 row 總和 與 輸出結果 ===
// 54
// 47
// 45
// 46
// 50
// 48
// 49
// 52
// 51
// 53
// ~~~

// ~~~
// 併發讀取的值 9
// 併發讀取的值 4
// 併發讀取的值 5
// 併發讀取的值 6
// 併發讀取的值 7
// 併發讀取的值 8
// 併發讀取的值 1
// 併發讀取的值 2
// 併發讀取的值 0
// 併發讀取的值 3
// ~~~
