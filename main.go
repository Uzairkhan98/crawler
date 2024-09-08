package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(argsWithProg) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of: ", argsWithProg[1])
	pages := make(map[string]int)
	baseURL, err := url.Parse(argsWithProg[1])
	if err != nil {
		fmt.Println("Invalid url: ", err)
		os.Exit(1)
	}
	cfg := &config{
		pages:              pages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
	}

	go cfg.crawlPage(baseURL.String())
	cfg.wg.Add(1)
	cfg.wg.Wait()

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
