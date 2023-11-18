package crawl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var (
	mu sync.Mutex
)

func getURLsFromHTML(htmlBody, baseURL string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}

	var urls []string

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if strings.HasPrefix(href, "/") {
				href = baseURL + href
			}

			if isValidURL(href) {
				urls = append(urls, href)
			}
		}
	})

	return urls
}

func normalizeURL(urlStr string) string {
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	pathNameLower := strings.ToLower(urlObj.Path)
	hostPath := urlObj.Hostname() + pathNameLower

	if len(hostPath) > 0 && strings.HasSuffix(hostPath, "/") {
		return hostPath[:len(hostPath)-1]
	}

	return hostPath
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func crawlPage(baseURL string, currentURL string, pages map[string]int, extPages map[string]interface{}) (map[string]int, map[string]interface{}) {
	baseURLObj, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return pages, extPages
	}

	currentURLObj, err := url.Parse(currentURL)
	if err != nil {
		fmt.Println("Error parsing current URL:", err)
		return pages, extPages
	}

	normalizedURL := normalizeURL(currentURL)

	if baseURLObj.Hostname() != currentURLObj.Hostname() {
		if count, ok := extPages[normalizedURL].(int); ok {
			extPages[normalizedURL] = count + 1
		} else if normalizedURL != "mailto:" && normalizedURL != "tel:" && normalizedURL != "javascript:void(0)" {
			extPages[normalizedURL] = 1
		}
		return pages, extPages
	}

	if count, ok := pages[normalizedURL]; ok {
		pages[normalizedURL] = count + 1
		return pages, extPages
	} else {
		pages[normalizedURL] = 1
	}

	fmt.Printf("actively crawling %s\n", currentURL)

	response, err := http.Get(currentURL)
	if err != nil {
		fmt.Printf("error in fetch: %s, on page: %s\n", err, currentURL)
		return pages, extPages
	}
	defer response.Body.Close()

	if response.StatusCode > 399 {
		fmt.Printf("error in fetch with status code: %d, on page: %s\n", response.StatusCode, currentURL)
		return pages, extPages
	}

	if response.ContentLength == 0 {
		fmt.Printf("empty html body found on page: %s\n", currentURL)
		return pages, extPages
	}

	contentType := response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		fmt.Printf("not an html page: %s\n", currentURL)
		return pages, extPages
	}

	htmlBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s, on page: %s\n", err, currentURL)
		return pages, extPages
	}

	urls := getURLsFromHTML(string(htmlBody), baseURL)

	for _, url := range urls {
		pages, extPages = crawlPage(baseURL, url, pages, extPages)
	}

	return pages, extPages
}
