package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseGlob("./templates/index.html")
	var shortener Shortener
	var url string
	var shortStr string

	if r.Method == "POST" {
		// TODO: Change this to response with json
		hostURL := os.Getenv("HOST_URL")
		r.ParseMultipartForm(512)

		if r.FormValue("url") != "" {
			url = r.FormValue("url")
			shortStr = shortener.shortString()
			saveURL(url, shortStr)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(hostURL + shortStr))
		}
		return
	}

	tmpl.Execute(w, nil)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	redirectURL := urldb[path].URL
	http.Redirect(w, r, redirectURL, 301)
}
