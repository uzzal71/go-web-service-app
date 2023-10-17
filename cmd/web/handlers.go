package main 

import (
	"net/http"
	"fmt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "</u></body></html>")
}

func (app *application) BookView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "View a single book")
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new book record form")
}