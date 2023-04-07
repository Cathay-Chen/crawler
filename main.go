package main

import (
	"bytes"
	"fmt"
	"github.com/Cathay-Chen/crawler/collect"
	"github.com/Cathay-Chen/crawler/proxy"
	"github.com/PuerkitoBio/goquery"
	"time"
)

func main() {
	proxyURLs := []string{"http://127.0.0.1:7890", "http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher error: ", err)
	}
	url := "https://www.google.com"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 3000 * time.Millisecond,
		Proxy:   p,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Println("Get error: ", err)
		return
	}
	fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println("NewDocumentFromReader error: ", err)
		return
	}

	doc.Find("div.news_li h2 a[target=_blank]").Each(func(i int, s *goquery.Selection) {
		// 获取匹配元素的文本
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
