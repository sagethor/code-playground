package main

import (
	"github.com/gocolly/colly/v2"
	//"time"
	"fmt"
	"log"
)

func main() {
	// initialize new colly Collector
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		// colly.AllowedDomains(),
		// Allow visiting the same page multiple times
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		colly.Async(true),
	);

	/*
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		// DomainGlob: "",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	});
	*/


	// add callbacks to Collector
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL);
	});

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err);
	});

	c.OnResponseHeaders(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL);
	});

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL);
	});

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"));
	});

	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text);
	});

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL);
	});


	c.Visit("https://old.reddit.com/r/popular/?count=100");
}
