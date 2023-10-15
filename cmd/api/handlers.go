package main 

/*
BODY='{"title": "The Black Soulston", "published":2001,"pages":107, "genres":["Fiction", "Mystery"], "rating":4.5}'
url -i -d "$BODY" localhost:4000/v1/books
BODY='{"title": "The Black Soulston", "published":2001,"pages":107, "genres":["Fiction", "Mystery"], "rating":3.5}'
curl -X PUT -d  "$BODY" localhost:4000/v1/books/12
{123 2023-10-14 01:08:20.669677485 +0600 +06 m=+3.199003110 The Black Soulston 2001 107 [Fiction Mystery] 3.5 1} 1 means true
*/
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
		books, err := app.models.Book.GetAll()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if err := app.writeJSON(w, http.StatusOK, envelope{"books": books}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		var input struct {
			Title    string   `json:"title"`
			Published int      `json:"published"`
			Pages    int      `json:"pages"`
			Genres   []string `json:"genres"`
			Rating   float32  `json:"rating"`
		}

		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		book := &data.Book {
			Title:		input.Title,
			Published:	input.Published,
			Pages:		input.Pages,
			Genres:		input.Genres,
			Rating:		input.Rating,
		}

		err = app.models.Books.Insert(book)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("v1/books/%d", book.ID))

		// Write the JSON response with a 201 Created status code and the Location header set.
		err = app.writeJSON(w, http.StatusCreated, envelope{"books": books}, headers);
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
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

	book, err := app.models.Books.Get(idInt)
	if err != nil {
		switch {
		case errors.is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest) // Use http.StatusBadRequest for a 400 Bad Request response
		return
	}

	book, err := app.models.Books.Get(idInt)
	if err != nil {
		switch {
		case errors.is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	
	var input struct {
		Title    *string   `json:"title"`
		Published *int      `json:"published"`
		Pages    *int      `json:"pages"`
		Genres   []string `json:"genres"`
		Rating   *float32  `json:"rating"`
	}


	err = app.readJSON(w, r, &input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.Published != nil {
		book.Published = *input.Published
	}

	if input.Pages != nil {
		book.Pages = *input.Pages
	}

	if len(input.Genres) > 0 {
		book.Genres = input.Genres
	}

	if input.Rating != nil {
		book.Rating = *input.Rating
	}

	err = app.models.Books.Update(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest) // Use http.StatusBadRequest for a 400 Bad Request response
		return
	}

	book, err := app.models.Books.Delete(idInt)
	if err != nil {
		switch {
		case errors.is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "book sucessfully deleted"}, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
