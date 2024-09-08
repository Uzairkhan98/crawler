package main

import (
	"fmt"
	"os"
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
	const maxConcurrency = 10

	cfg, err := configure(argsWithProg[1], maxConcurrency)

	if err != nil {
		fmt.Printf("error configuring: %v\n", err)
		os.Exit(1)
	}

	go cfg.crawlPage(argsWithProg[1])
	cfg.wg.Add(1)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
