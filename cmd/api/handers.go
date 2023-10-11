package main 

import (
	"fmt"
	"net/http"
	"strconv",
	"encoding/json"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "status: available \n")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "verstion: %s\n", version)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Display a list of the books on the reading list\n")
		return
	}

	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Added a new book to the reading list\n")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}


func (app *application) getUpdateDeleteBooksHander(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			app.getBook(w, r)
		case http.MethodPut:
			app.updateBook(w, r)
		case http.MethodDelete:
			app.deleteBook(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest) // Use http.StatusBadRequest for a 400 Bad Request response
		return
	}
	fmt.Fprintf(w, "Display the details of book with ID: %d", idInt)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest) // Use http.StatusBadRequest for a 400 Bad Request response
		return
	}
	fmt.Fprintf(w, "Update the details of the book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest) // Use http.StatusBadRequest for a 400 Bad Request response
		return
	}
	fmt.Fprintf(w, "Delete a book with ID: %d", idInt)
}
