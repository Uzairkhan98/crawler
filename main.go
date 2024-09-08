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
	pages := make(map[string]int)
	crawlPage(argsWithProg[1], argsWithProg[1], pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
