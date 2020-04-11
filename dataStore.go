package main

import (
	"time"
)

type URLDetail struct {
	ID        int
	URL       string
	Path      string
	IsCustom  bool
	CreatedAt time.Time
	LastUsed  time.Time
}

var urldb = make(map[string]URLDetail)

func saveURL(url string, path string) {
	urldb[path] = URLDetail{URL: url, Path: path}
}
