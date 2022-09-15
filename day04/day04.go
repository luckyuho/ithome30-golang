package main

import "fmt"

func main() {
	forExample()
	mapCombineStructIteration()
	funcExample2()
}

func forExample() {

	// 完整寫法
	// fmt.Printf()與fmt.Println()相似，但是其中有一些差異
	// 可以參考下面參考來源提供的第一個連結
	fmt.Printf("完整寫法: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i) // 0 1 2 3 4 5 6 7 8 9
	}

	// 缺少初始值設定的寫法
	fmt.Printf("缺少初始: ")
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("%d ", j) // 0 1 2 3 4 5 6 7 8 9
	}

	// 模擬while的方式
	fmt.Printf("模擬寫法: ")
	k := 0
	for k < 10 {
		fmt.Printf("%d ", k) // 0 1 2 3 4 5 6 7 8 9
		k++
	}
}

type Nominate struct {
	President     string
	VicePresident string
}

func mapCombineStructIteration() {
	electionCandidates := map[int]Nominate{
		1: {"宋楚瑜", "余湘"},
		2: {"韓國瑜", "張善政"},
		3: {"蔡英文", "賴清德"},
	}
	// for i, v := range electionCandidates {
	// 	fmt.Printf("號次: %d, 總統候選人: %s, 副總統候選人: %s\n", i, v.President, v.VicePresident)
	// }
	output := canvassing(electionCandidates, "總統", 2020)
	fmt.Println(output)
}

func funcExample2() {
	electionCandidates := map[int]Nominate{
		1: {"黃晶", "林欣汝"},
		2: {"柯灃隆", "蔡承諺"},
	}
	output := canvassing(electionCandidates, "成大學生會正副會長", 2020)
	fmt.Println(output)
}

func canvassing(candidates map[int]Nominate, title string, year int) string {
	for i, v := range candidates {
		fmt.Printf("%d 請大家多多支持🙏🙏🙏 %d號%s候選人～%s和%s，凍蒜！凍蒜！\n", year, i, title, v.President, v.VicePresident)
	}

	totalCandidates := len(candidates)
	result := fmt.Sprintf("%d 總共有 %d 組%s候選人", year, totalCandidates, title)
	return result
}
