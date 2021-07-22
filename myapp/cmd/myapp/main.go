package main

import (
	"log"
	"os"

	"github.com/benbjohnson/myapp/http"
	"github.com/benbjohnson/myapp/postgres"
)

func main() {
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}

	var h http.Handler
	h.UserService = us
}
