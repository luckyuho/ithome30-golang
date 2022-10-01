package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const url = "https://ithelp.ithome.com.tw/users/20150797/ironman/5271?page=2"

// 從網路上讀取資料
func Fetcher(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request error %s", err)
	}

	// 有些 url 需要有 User-Agent 的資訊，否則會出現 403 的問題，鐵人30 就需要這個東西
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("client error %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// 從檔案讀取資料
func ReadFile(fileName string) ([]byte, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("read file error: %s", err)
	}

	return file, err
}

// 正則表達式
var HotQuestionRe = `class=\"qa-list__title-link\">\s*([^<]*)`

// 檔案路徑
var filePath = "day20/itTitle.html"

func main() {
	// 從網路獲取文章
	html, err := Fetcher(url)

	// 從文檔讀取文章
	// html, err := ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}

	//解析正则表達式，如果成功返回解释器
	reg := regexp.MustCompile(HotQuestionRe) // \t表空格，[^\t]表示除了空格外，其他都可
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	// 文章匹配正則表達式
	result := reg.FindAllSubmatch(html, -1)
	var s []string
	for _, v := range result {
		trimV := strings.TrimSpace(string(v[1]))
		s = append(s, trimV)
	}

	// 輸出結果
	for i := range s {
		fmt.Println(s[i])
	}
}
