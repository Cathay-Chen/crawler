package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Print("fetch url error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("read content failed: %v", err)
		return
	}

	// fmt.Println("Body: ", string(body))

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links! \n", numLinks)

	numLinks = bytes.Count(body, []byte("<a"))
	fmt.Printf("homepage has %d links! \n", numLinks)

	exist := strings.Contains(string(body), "美国")
	fmt.Printf("Is there a word for 美国: %v\n", exist)

	exist = bytes.Contains(body, []byte("人民"))
	fmt.Printf("Is there a word for 人民: %v\n", exist)
}
