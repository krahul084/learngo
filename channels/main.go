package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.bing.com",
		"http://www.golang.org",
		"http://duckduckgo.com",
	}
	c := make(chan string)
	for _, link := range links {
		go getStatus(link, c)

	}
	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			go getStatus(l, c)
		}(l)
	}

}

func getStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Not Up")
		c <- link
		return
	}
	fmt.Println(link, "Up")
	c <- link
}
