package main 

import (
	"net/http"
	"fmt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The home page")
}

func (app *application) BookView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "View a single book")
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The home page")
}