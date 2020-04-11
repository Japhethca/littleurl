package main

import (
	"github.com/google/uuid"
)

// Shortener shortens or generates random string of specified length
type Shortener struct {
	length int // string length
	str    string
}

func (s *Shortener) shortString() string {
	if s.length == 0 {
		s.setLength(6)
	}

	if s.str == "" {
		s.setString(uuid.New().String())
	}

	if len(s.str) < s.length {
		return s.str
	}
	return s.str[:s.length]
}

func (s *Shortener) String() string {
	return s.shortString()
}

func (s *Shortener) setString(str string) {
	if len(str) > 0 {
		s.str = str
	}
}

func (s *Shortener) setLength(l int) {
	if l > 0 {
		s.length = l
	}
}
