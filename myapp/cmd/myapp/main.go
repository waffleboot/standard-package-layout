package main

import (
	"log"
	"net/http"
	"os"

	myhttp "github.com/benbjohnson/myapp/http"
	"github.com/benbjohnson/myapp/postgres"
)

func main() {
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}

	var h myhttp.Handler
	h.UserService = us
	if err := http.ListenAndServe(":8080", &h); err != nil {
		log.Fatal(err)
	}
}
