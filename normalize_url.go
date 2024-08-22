package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %w", err)
	}
	normalizedUrl := parsedUrl.Host + parsedUrl.Path
	normalizedUrl = strings.ToLower(normalizedUrl)
	normalizedUrl = strings.TrimSuffix(normalizedUrl, "/")
	return normalizedUrl, nil
}
