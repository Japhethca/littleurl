package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type DBActions interface {
	saveURL(url string, path string, options urlOptions) (URLDetail, error)
	getURLByPath(path string) (URLDetail, error)
	getPathByURL(url string) (URLDetail, error)
	updateUrlPath(url, path string, options urlOptions) (URLDetail, error)
}

type handler struct {
	db DBActions
}

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseGlob("./templates/index.html")
	var shortener Shortener

	if r.Method == "POST" {
		// TODO: Change this to response with json
		hostURL := os.Getenv("HOST_URL")
		r.ParseMultipartForm(512)

		if r.FormValue("url") == "" {
			return
		}

		var ud URLDetail
		url := r.FormValue("url")
		customPath := r.FormValue("custom-path")
		ud, _ = h.db.getPathByURL(url)
		if ud.URL == "" {
			var err error
			urlopt := urlOptions{isCustom: true}
			pathStr := customPath

			if pathStr == "" {
				pathStr = shortener.shortString()
				urlopt.isCustom = false
			}

			ud, err = h.db.saveURL(url, pathStr, urlopt)
			if err != nil {
				log.Print(err)
				msg := []byte("Something went wrong, we were unable to create a short url.")
				respond(w, msg, http.StatusInternalServerError)
				return
			}
		}

		if ud.IsCustom == false && customPath != "" {
			var err error
			ud, err = h.db.updateUrlPath(url, customPath, urlOptions{isCustom: true})
			if err != nil {
				log.Fatal(err) // TODO: handle server errors
			}
		}
		respond(w, []byte(hostURL+ud.Path), http.StatusOK)
		return
	}

	tmpl.Execute(w, nil)
}

func (h *handler) redirect(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	ud, err := h.db.getURLByPath(path)
	if err != nil {
		http.NotFound(w, r) // TODO: create template for 404 eror page
		return
	}
	http.Redirect(w, r, ud.URL, 301)
}

func respond(w http.ResponseWriter, data []byte, statuscode int) {
	w.WriteHeader(statuscode)
	w.Write(data)
}
