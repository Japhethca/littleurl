package main

import "testing"

func TestStringShortener(t *testing.T) {
	shortener := Shortener{}

	t.Run("Test returns default generated string length", func(t *testing.T) {
		str := shortener.shortString()
		if len(str) != shortener.length {
			t.Errorf("Generated string '%s' should have a default length %d", str, shortener.length)
		}
	})

	t.Run("Test returns generated string with given length", func(t *testing.T) {
		el := 5
		shortener.setLength(el)
		str := shortener.shortString()
		if len(str) != el {
			t.Errorf("Generated string '%s' should have length of %d", str, el)
		}
	})
}
