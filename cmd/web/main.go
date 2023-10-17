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
}