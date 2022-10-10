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

// 總和 1 個 row 的值
func sum1Row(arr1Dim []int, c chan int) {
	sum := 0
	for _, v := range arr1Dim {
		sum += v
	}
	c <- sum // 總和後傳給 channel
}

// 總和 1 個 row 的值 與 顯示
func sum1RowAndShow(arr [][]int, m, n int) {
	c := make(chan int)

	// 併發讀取與寫入
	for i := 0; i < m; i++ {
		go sum1Row(arr[i], c)
	}

	// 這時候只有 c，就可以正常併發讀取
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", <-c)
		fmt.Println()
	}

}

func main() {
	m, n := 10, 10
	arr := create2DimArray(m, n)

	showTitle("=== 顯示 2 維 arr 內容物 ===")
	show2DimArray(arr)

	showTitle("=== goroutine 對 2 維每 row 總和 與 輸出結果 ===")
	sum1RowAndShow(arr, m, n)

}

func showTitle(s string) {
	fmt.Println(s)
}
