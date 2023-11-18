# Go URL Crawler

## Overview

The Go URL Crawler is a lightweight web crawler written in GoLang. It is designed to recursively crawl web pages, extract URLs, and provide insights into the structure of a website.

## Features

- **URL Normalization:** Normalize URLs by converting pathnames to lowercase and handling trailing slashes.
- **HTTP Fetching:** Retrieves HTML content from web pages using the standard `net/http` package.
- **Error Handling:** Provides detailed error messages for fetch failures, HTTP errors, and HTML parsing errors.
- **Content Type Validation:** Ensures that the fetched content is of type "text/html" before proceeding with URL extraction.
- **Recursive Crawling:** Continues to crawl extracted URLs to build a comprehensive map of the website structure.

## Installation

- git clone https://github.com/mrmissong/url-crawler-go.git
- go get github.com/mrmissong/url-crawler-go
- go run main.go

## Dependencies

This project relies on the following external packages:

- **github.com/PuerkitoBio/goquery:** A package for HTML parsing.
- **net/http:** The standard HTTP package for making HTTP requests.
- **strings:** The standard package for string manipulation.
- **sync:** The standard package for synchronization (mutex).

## Contributing

Contributions are welcome! If you have ideas for improvements or find any issues, please feel free to open an issue or create a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

