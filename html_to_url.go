package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func isValidURL(testURL string) bool {
	parsedURL, err := url.ParseRequestURI(testURL)
	if err != nil {
		return false
	}

	// Ensure the scheme is either HTTP or HTTPS and the host is present
	if (parsedURL.Scheme == "http" || parsedURL.Scheme == "https") && parsedURL.Host != "" {
		return true
	}

	return false
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)

	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					val := a.Val
					if strings.Contains(val, "mailto:") {
						continue
					}
					if !strings.HasPrefix(val, "http") {
						val = baseURL.String() + val
					}
					if isValidURL(val) {
						res = append(res, val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return res, nil
}
