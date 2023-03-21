package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urlList := []string{
		"http://google.com",
		"http://reddit.com",
		"http://stackoverflow.com",
		"http://golang.org",
		// "https://dev.zed",
		// "https://daredfjaghed.comd",
	}

	c := make(chan string)

	for _, url := range urlList {
		go checkUrl(url, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkUrl(l, c)
		}(l)
	}

}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "\t\t..... ERR")
		c <- url
		return
	}

	fmt.Println(url, "\t\t..... OK")
	c <- url

}
