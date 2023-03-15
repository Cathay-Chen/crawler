package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

//var headerRe = regexp.MustCompile(`<div class="small_cardcontent__BTALp"[\s\S]*?<h2>([\s\S]*?)</h2>`)

func main() {
	url := "https://www.thepaper.cn/"
	body, err := Fetch(url)

	if err != nil {
		fmt.Println("read content failed: %v", err)
		return
	}

	//matches := headerRe.FindAllSubmatch(body, -1)
	//
	//for _, m := range matches {
	//	fmt.Println("fetch card news:", string(m[1]))
	//}

	//doc, err := htmlquery.Parse(bytes.NewReader(body))
	//
	//if err != nil {
	//	fmt.Println("htmlquery.Parse failed: %v", err)
	//}
	//
	//nodes := htmlquery.Find(doc, `//div[@class='small_cardcontent__BTALp']//h2`)
	//
	//for _, node := range nodes {
	//	fmt.Println("fetch card news:", node.FirstChild.Data)
	//}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Println("goquery.NewDocumentFromReader failed: %v", err)
	}

	doc.Find("div.small_cardcontent__BTALp h2").Each(func(i int, s *goquery.Selection) {
		fmt.Println("fetch card news:", i, s.Text())
	})
}

// Fetch 用于获取网页内容
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error status code: %v", resp.StatusCode)
	}

	// 读取网页内容，判断编码，转换成 utf8，再转换成 []byte
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// DetermineEncoding 用于判断网页的编码
func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	// 读取前 1024 个字节
	bytes, err := r.Peek(1024)

	// 如果读取出错，就返回 utf8
	if err != nil {
		fmt.Println("fetch error: %v", err)
		return unicode.UTF8
	}

	// 判断编码
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
