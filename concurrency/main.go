package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("Type is : %v \n", ctype)
}

func main() {
	sites := []string{
		"https://mail.google.com/",
		"https://httpbin.org/xml",
		"https://golang.org",
		"https://api.github.com",
	}
	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)
		go func(site string) {
			returnType(site)
			wg.Done()
		}(site)

	}

	wg.Wait()

}
