package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortUrl := shorten(r.URL.Path)
		fmt.Fprintf(w, "Your shortened url is %s.\n", shortUrl)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
