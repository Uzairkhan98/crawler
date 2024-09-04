package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func normalizeURL(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	u.Host = strings.ToLower(u.Host)
	u.Path = path.Clean(u.Path)

	res := path.Clean(u.Host)

	if u.Path != "/" {
		res = u.Host + u.Path
	}

	return res, nil
}
