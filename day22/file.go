package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 讀取全部
func ReadFileAll(filePath string) []byte {
	file, err := os.Open(filePath)
	if err == nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	return data
}

// 逐行讀取
func ReadFileLine(filePath string) []byte {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	lines := make([]byte, 0)
	r := bufio.NewReader(fi)
	for {
		line, _, c := r.ReadLine()
		if c == io.EOF {
			break
		}
		line = []byte(string(line) + "\n") // 為了與其他結果一樣而多加的一行
		lines = append(lines, line...)
	}

	return lines
}

// 讀取字元
func ReadFileChar(filePath string) []byte {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次讀取1024個字元
	for {
		n, err := r.Read(buf)
		if err == io.EOF || err != nil {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}

	return chunks
}

func main() {
	// filePath := "day22/itTitle.html"

	// 讀取
	// 將整份文件獨進記憶體
	// data := ReadFileAll(filePath)
	// 讀字元
	// data := ReadFileChar(filePath)
	// 讀行
	// data := ReadFileLine(filePath)
	// fmt.Printf("%s", data)

	// 寫入
	for i := 0; i < 2; i++ {
		WriteFileCover("day22/dat1.yaml")
		WriteFileContinue("day22/dat2.yaml")
	}
}

var str = []byte("hello\ngo\n")

// 寫入文檔
func WriteFileCover(fileName string) {
	err := os.WriteFile(fileName, str, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteFileContinue(fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	_, err = f.WriteString(string(str))
	if err != nil {
		fmt.Println(err)
	}
}
