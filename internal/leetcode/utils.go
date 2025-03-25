package leetcode

import (
	"strings"
)

// Validate URL
func checkURLPrefix(url string) error {
	if strings.HasPrefix(url, "https://leetcode.com/problems/") {
		return nil
	} else {
		return ErrNotCorrectPrefix
	}
}

// Extract problem title from url
func getProblemTitle(url string) (string, error) {
	urlParts := strings.Split(url, "/")
	if len(urlParts) > 4 {
		return urlParts[4], nil
	} else {
		return "", ErrIncorrectURL
	}

}
