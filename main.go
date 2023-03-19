package main

import (
	"bytes"
	"fmt"
	"github.com/Cathay-Chen/crawler/collect"
	"github.com/PuerkitoBio/goquery"
	"time"
)

//var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	url := "https://book.douban.com/subject/1007305/" // 红楼梦
	var f collect.Fetcher = collect.BowerFetcher{
		Timeout: 3 * time.Second,
	}

	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read cotent failed: %v\n", err)
		return
	}

	// 加载HTML文档
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("load html content failed: %v\n", err)
	}
	//fmt.Println(doc.Html())
	doc.Find("#wrapper h1 span").Each(func(i int, selection *goquery.Selection) {
		title := selection.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

}
