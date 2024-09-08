package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	cfg.concurrencyControl <- struct{}{}
	if pageLimitReached := cfg.pageLimitReached(); pageLimitReached {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL: %v", err)
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}
	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing URL: %v", err)
	}
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error fetching HTML for URL: %v, %v", rawCurrentURL, err)
	}
	fmt.Println(html)
	newURLs, err := getURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		fmt.Printf("error extracting URLs from HTML: %v", err)
	}

	for _, newURL := range newURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(newURL)
	}
}
