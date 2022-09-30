package main

import (
	"fmt"
	"regexp"
)

const ComparedSentence = "abc@ade.com abd@bcd.com"

func main() {
	//解析正则表達式，如果成功返回解释器
	reg := regexp.MustCompile(`a[^\t]+.com`) // \t表空格，[^\t]表示除了空格外
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	result := reg.FindAllStringSubmatch(ComparedSentence, -1)
	fmt.Println("result = ", result)
}
