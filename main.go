package main

import (
	crawl "crawler/crawler"
	"fmt"
)

func main() {
	fmt.Print("Which URL do you want to crawl? ")
	var urlName string
	fmt.Scanln(&urlName)

	crawl.CrawlURL(urlName)
}
