package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {

}

func main() {
	godotenv.Load()
	db, err := sql.Open(os.Getenv("DB_DIALECT"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	dbm := &dbManager{db}

	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	h := &handler{dbm}
	r.HandleFunc("/", h.home)
	r.HandleFunc("/{urlpath}", h.redirect)
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
