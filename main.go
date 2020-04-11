package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fs := http.FileServer(http.Dir("./static/"))

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/", home)
	r.HandleFunc("/{urlpath}", redirect)
	http.ListenAndServe(":8080", r)
}
