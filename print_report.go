package main

import (
	"fmt"
	"net/url"
	"sort"
)

type Page struct {
	URL   string
	Count int
}

// sortPages sorts the map into a slice of Page structs, ordered by link count,
// and alphabetically for pages with the same count.
func sortPages(pages map[string]int) []Page {
	// Convert the map to a slice of Page structs
	pageList := make([]Page, 0, len(pages))
	for url, count := range pages {
		pageList = append(pageList, Page{
			URL:   url,
			Count: count,
		})
	}

	// Sort the slice: first by Count in descending order, then by URL alphabetically
	sort.Slice(pageList, func(i, j int) bool {
		if pageList[i].Count == pageList[j].Count {
			return pageList[i].URL < pageList[j].URL // Sort alphabetically by URL if count is the same
		}
		return pageList[i].Count > pageList[j].Count // Sort by count (descending)
	})

	return pageList
}

func printReport(pages map[string]int, baseURL string) {
	url, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("Error parsing base URL: %v\n", err)
		return
	}
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")
	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, url.Scheme+"://"+page.URL)
	}
}
