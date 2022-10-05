package main

import "it/day24/routers"

func main() {
	r := routers.InitRouter()

	// 定義在哪個 port 上執行
	r.Run(":8080")
}
