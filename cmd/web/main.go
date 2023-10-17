package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
	"readinglist.uzzal.io/internal/models"
)

const version = "10.0"

type config struct {
	port int
	env  string
	dsn string
}

type application struct {
	config config
	logger *log.Logger
	readinglist *models.ReadinglistModel
}

func main() {
	var cfg config
	var db *sql.DB

	addr := flag.String("addr", ":8080", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:8080/v1/books", "Endpoint for the readinglist web service")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)") // Use StringVar here
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("database connection pool established")

	app := &application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint},
	}

	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
	}

	log.Printf("Staring the server on %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}