package main

import (
	"flag"
	"logs"
	"net/http"
)

type application struct {

}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")
	
	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes()
	}

	log.Printf("Staring the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}