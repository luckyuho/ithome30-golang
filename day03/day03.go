package main

import "fmt"

func main() {
	mapShow()
	structShow()
	mapCombineStruct()
}

func mapShow() {
	presidentialCandidate := map[int]string{
		1: "宋楚瑜",
		2: "韓國瑜",
		3: "蔡英文",
	}
	vicePresidentialCandidate := map[string]string{
		"宋楚瑜": "余湘",
		"韓國瑜": "張善政",
		"蔡英文": "賴清德",
	}

	fmt.Println("編號與總統候選人:")
	for k, v := range presidentialCandidate {
		fmt.Println(k, v)
	}
	fmt.Println("總統與副總統候選人:")
	for k, v := range vicePresidentialCandidate {
		fmt.Println(k, v)
	}
}

type Presidential struct {
	Id   int
	Name string
}

type Nominate struct {
	President     string
	VicePresident string
}

type ElectionCandidates struct {
	Id     int
	Couple Nominate
}

func structShow() {
	president := []Presidential{
		{Id: 1, Name: "宋楚瑜"},
		{Id: 2, Name: "韓國瑜"},
		{Id: 3, Name: "蔡英文"}}
	fmt.Println(president)
	electionCandidates := []ElectionCandidates{
		{Id: 1, Couple: Nominate{President: "宋楚瑜", VicePresident: "余湘"}},
		{Id: 2, Couple: Nominate{President: "韓國瑜", VicePresident: "張善政"}},
		{Id: 3, Couple: Nominate{President: "蔡英文", VicePresident: "賴清德"}},
	}
	fmt.Println(electionCandidates)
}

func mapCombineStruct() {
	electionCandidates := map[int]Nominate{
		1: {"宋楚瑜", "余湘"},
		2: {"韓國瑜", "張善政"},
		3: {"蔡英文", "賴清德"},
	}
	fmt.Println(electionCandidates)
}
