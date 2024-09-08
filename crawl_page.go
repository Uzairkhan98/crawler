package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()
	cfg.concurrencyControl <- struct{}{}

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

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	isFirst, cfg.pages[normalizedURL] = cfg.pages[normalizedURL] == 0, cfg.pages[normalizedURL]+1
	return isFirst
}
