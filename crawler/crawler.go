package crawl

import (
	"fmt"
	"strings"
)

func CrawlURL(urlName string) {
	trimmedURL := strings.TrimSpace(urlName)

	if strings.Contains(trimmedURL, " ") {
		fmt.Println("Please try again and enter only one URL")
		return
	}

	if strings.HasPrefix(trimmedURL, "http://") || strings.HasPrefix(trimmedURL, "https://") {
		fmt.Printf("Crawling %s...\n", trimmedURL)

		pages := make(map[string]int)
		extPages := make(map[string]interface{})

		pages, extPages = crawlPage(trimmedURL, trimmedURL, pages, extPages)

		// Print report
		fmt.Println("Pages:")
		for page, count := range pages {
			fmt.Printf("%s: %d\n", page, count)
		}

		fmt.Println("External Pages:")
		for page, count := range extPages {
			fmt.Printf("%s: %d\n", page, count.(int))
		}
	} else {
		fmt.Println("Please try again and enter a valid URL that starts with http:// or https://")
	}
}
