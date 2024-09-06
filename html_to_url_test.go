package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestHTMLToURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      []string
		errorContains string
		inputBody     string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "multiple absolute URLs",
			inputURL: "https://example.com",
			inputBody: `
		<html>
			<body>
				<a href="https://site1.com/page1">Site1</a>
				<a href="https://site2.com/page2">Site2</a>
				<a href="https://site3.com/page3">Site3</a>
			</body>
		</html>
		`,
			expected: []string{"https://site1.com/page1", "https://site2.com/page2", "https://site3.com/page3"},
		},
		{
			name:     "mixed relative and absolute URLs",
			inputURL: "https://example.com",
			inputBody: `
		<html>
			<body>
				<a href="/internal/page1">Internal1</a>
				<a href="/internal/page2">Internal2</a>
				<a href="https://external.com/page3">External</a>
			</body>
		</html>
		`,
			expected: []string{"https://example.com/internal/page1", "https://example.com/internal/page2", "https://external.com/page3"},
		},
		{
			name:     "no links in page",
			inputURL: "https://example.com",
			inputBody: `
		<html>
			<body>
				<p>Welcome to the example page. No links here!</p>
			</body>
		</html>
		`,
			expected: []string{},
		},
		{
			name:     "malformed URLs",
			inputURL: "https://example.com",
			inputBody: `
		<html>
			<body>
				<a href="https://validsite.com/">Valid Site</a>
				<a href="javascript:void(0)">Invalid JS URL</a>
				<a href="mailto:someone@example.com">Email Link</a>
			</body>
		</html>
		`,
			expected: []string{"https://validsite.com/"},
		},
		{
			name:     "links with hashes and query params",
			inputURL: "https://example.com",
			inputBody: `
		<html>
			<body>
				<a href="/page1?query=test">Query Link</a>
				<a href="/page2#section">Hash Link</a>
				<a href="https://external.com/page3?data=info#anchor">External with Params</a>
			</body>
		</html>
		`,
			expected: []string{"https://example.com/page1?query=test", "https://example.com/page2#section", "https://external.com/page3?data=info#anchor"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
