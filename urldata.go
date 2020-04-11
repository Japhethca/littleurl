package main

import (
	"time"
)

type urlOptions struct {
	isCustom bool
}

type URLDetail struct {
	URL       string    `json:"url"`
	Path      string    `json:"path"`
	IsCustom  bool      `json:"is_custom"`
	CreatedAt time.Time `json:"createdAt"`
	LastUsed  time.Time `json:"lastUsed"`
}
