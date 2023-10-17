package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct {

}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")

	app := &application{}
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
	}

	log.Printf("Staring the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}