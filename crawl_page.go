package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL: %v", err)
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL: %v", err)
	}
	if baseURL.Host != currentURL.Host {
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL: %v", err)
	}
	val, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL] = val + 1
		return
	} else {
		pages[normalizedURL] = 1
	}
	html, err := getHTML(currentURL.Scheme + "://" + normalizedURL)
	if err != nil {
		fmt.Printf("error fetching HTML for URL: %v, %v", normalizedURL, err)
	}
	fmt.Println(html)
	newURLs, err := getURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		fmt.Printf("error extracting URLs from HTML: %v", err)
	}

	for _, newURL := range newURLs {
		crawlPage(rawBaseURL, newURL, pages)
	}
}
