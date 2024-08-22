package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, int, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", 0, fmt.Errorf("got Network error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", res.StatusCode, fmt.Errorf("got HTTP error: %s", res.Status)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", 0, fmt.Errorf("got non-HTML response: %s", contentType)
	}

	htmlBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", 0, fmt.Errorf("couldn't read response body: %v", err)
	}

	htmlBody := string(htmlBodyBytes)

	return htmlBody, res.StatusCode, nil
}
