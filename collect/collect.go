package collect

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"time"
)

type Fetcher interface {
	Get(url string) ([]byte, error)
}

type BaseFetcher struct {
}

func (BaseFetcher) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	// 转换编码
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

type BowerFetcher struct {
	Timeout time.Duration
}

// 模拟浏览器访问
func (f BowerFetcher) Get(url string) ([]byte, error) {
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		Dial: (&net.Dialer{
	//			Timeout:   30 * time.Second,
	//			KeepAlive: 30 * time.Second,
	//		}).Dial,
	//		TLSHandshakeTimeout:   10 * time.Second,
	//		ResponseHeaderTimeout: 10 * time.Second,
	//		ExpectContinueTimeout: 1 * time.Second,
	//	},
	//}
	client := &http.Client{
		Timeout: f.Timeout,
	}
	// 构造请求
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("get url failed: %v", err)
	}

	// 设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 读取网页内容，判断编码，转换成 utf8，再转换成 []byte
	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

// DetermineEncoding 用于判断网页的编码
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	// 读取前 1024 个字节
	bytes, err := r.Peek(1024)

	// 如果读取出错，就返回 utf8
	if err != nil {
		fmt.Printf("fetch error: %v\n", err)
		return unicode.UTF8
	}

	// 判断编码
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
