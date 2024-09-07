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
	html, err := getHTML(argsWithProg[1])
	if err != nil {
		fmt.Printf("error getting HTML: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(html)
}
