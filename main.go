package main

import (
	"fmt"
	"net/http"
	"time"
)

func getLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link

		return
	}
	fmt.Println(link, "is up")
	c <- link
}
func main() {
	links := []string{
		"http://linked.com",
		"http://google.com",
		"http://facebook.com",
	}
	c := make(chan string)
	for _, link := range links {
		go getLink(link, c)
	}
	for l := range c {
		go func(lin string) {
			time.Sleep(3 * time.Second)
			getLink(lin, c)
		}(l)

	}
}
