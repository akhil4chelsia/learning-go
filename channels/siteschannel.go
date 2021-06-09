package main

import (
	"fmt"
	"net/http"
)

func main() {

	sites := []string{
		"https://mail.google.com/",
		"https://httpbin.org/xml",
		"https://golang.org",
		"https://api.github.com",
	}
	ch := make(chan string)

	for _, url := range sites {
		go returnType(ch, url)
	}

	for val := range ch {
		fmt.Println(val)
	}
}

func returnType(ch chan string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	ch <- fmt.Sprintf("Type is : %v \n", ctype)
}
