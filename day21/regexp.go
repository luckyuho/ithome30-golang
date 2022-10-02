package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {

	collector := colly.NewCollector()

	// // 抓標籤Tag
	// c.OnHTML("title", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	// // 抓屬性值 AttrVal
	// collector.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Attr("content")) // 抓此Tag中的name屬性 來找出此Tag，再印此Tag中的content屬性
	// })

	// 抓 Class 名稱
	collector.OnHTML(".qa-list__title-link", func(e *colly.HTMLElement) {
		trimText := strings.TrimSpace(e.Text)
		fmt.Println(trimText)
	})

	// 請求期間發生錯誤處理
	collector.OnError(func(response *colly.Response, err error) {
		fmt.Println("error :", err)
	})

	// 有些 url 需要有 User-Agent 的資訊，否則會出現 403 的問題，鐵人30 就需要這個東西
	collector.OnRequest(func(req *colly.Request) {
		req.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	collector.Visit("https://ithelp.ithome.com.tw/users/20150797/ironman/5271?page=2")
}
