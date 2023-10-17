package main 

impport "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/book/view", app.BookView)
	mux.HandleFunc("/book/create", app.bookCreate)

	return mux
}