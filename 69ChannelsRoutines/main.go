package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(time.Second)
	_, err := http.Get(link)
	if err != nil {

		fmt.Println("check failed: ", link)
		c <- link
		return
	}
	fmt.Println("check ok:", link)
	c <- link
}
