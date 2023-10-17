package main

import (
	"flag"
	"log"
	"net/http"

	"readinglist.uzzal.io/internal/models"
)

type application struct {
	readinglist *models.ReadinglistModel
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:8080/v1/books", "Endpoint for the readinglist web service")
	
	app := &application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint},
	}

	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
	}

	log.Printf("Staring the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}