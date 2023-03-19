package main

import (
	"fmt"
	"net/http"
)

func main() {

	urlList := []string{
		"http://google.com",
		"http://reddit.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"https://dev.zed",
		"https://daredfjaghed.comd",
	}

	for _, url := range urlList {
		checkUrl(url)
	}

}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "......ERR")
		return
	}

	fmt.Println(url, "......OK")
}
