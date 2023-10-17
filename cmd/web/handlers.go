package main 

import (
	"net/http"
	"fmt"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	books, err := app.readinglist.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<html><head><title>Reading List</title></head><body><h1>Reading List</h1><ul>")
	for _, book := range *books {
		fmt.Fprintf(w, "<li> %s (%d)</li>", book.Title, book.Pages)
	}
	fmt.Fprintln(w, "</u></body></html>")
}

func (app *application) BookView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "View a single book")
}

func (app *application) bookCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new book record form")
}