package main 

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"time"
	"readinglist.uzzal.io/internal/data"
)


func (app *application) healthcheck(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}

	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		books := []data.Book{
			{
				ID:			1,
				CreatedAt:	time.Now(),
				Title:		"The Darkening of Tristram",
				Published:	1998,
				Generes:	[]string{"Fiction", "Thriller"},
				Rating:		4.5,
				Version:	1,
			},
			{
				ID:			2,
				CreatedAt:	time.Now(),
				Title:		"The Legecy of Deckar Cain",
				Published:	2007,
				Generes:	[]string{"Fiction", "Adventure"},
				Rating:		4.9,
				Version:	1,
			}
		}

		js, err := json.Marshal(books)
		if err != nil{
			http.Error(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		js = append(js, '\n')
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
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
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	book := data.Book{
		ID: idInt,
		CreatedAt: time.Now(),
		Title: "Echoes in the Darkness",
		Published: 2019,
		Pages: 300,
		Genres: []string{"Fiction", "Thriller"},
		Rating: 4.5,
		Version: 1,
	}

	js, err := json.Marshal(book)

	if err _= nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
