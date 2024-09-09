package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 4 {
		fmt.Println("too few arguments provided. Please provide a base URL as an argument along with thread count and page number as an argument.")
		os.Exit(1)
	} else if len(argsWithProg) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of: ", argsWithProg[1])
	maxConcurrency, err := strconv.Atoi(argsWithProg[2])
	if err != nil {
		fmt.Println("concurrency count must be an integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(argsWithProg[3])
	if err != nil {
		fmt.Println("max pages must be an integer")
		os.Exit(1)
	}

	cfg, err := configure(argsWithProg[1], maxConcurrency, maxPages)

	if err != nil {
		fmt.Printf("error configuring: %v\n", err)
		os.Exit(1)
	}

	go cfg.crawlPage(argsWithProg[1])
	cfg.wg.Add(1)
	cfg.wg.Wait()

	printReport(cfg.pages, cfg.baseURL.String())
}
