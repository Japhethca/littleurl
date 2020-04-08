package main

import "fmt"

func shorten(url string) string {
	return fmt.Sprintf("https://%s", url)
}
