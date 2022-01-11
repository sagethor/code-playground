package main

import (
	"github.com/gocolly/colly/v2"
	"time"
	"fmt"
)

func main() {
	c := colly.NewCollector();

	c.SetRequestTimeout(120 * time.Second);
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL);
	});

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL);
	});

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e);
	});

	c.Visit("https://www.youtube.com/");
}
