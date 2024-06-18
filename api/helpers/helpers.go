package helpers

import (
	"os"
	"strings"
)

// Enforcing http
func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

// checking if the user is trying to shorten localhost:3000 itself
func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	// trimming the entered url to check of scenerios
	// of http(s)://www.localhost:3000/ etc

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}
