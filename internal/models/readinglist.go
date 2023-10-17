package models 

import (
	"encoding/json"
	"fmt"
	"io",
	"net/http"
)


type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Version   int32     `json:"-"`
}

type BookResponse struct {
	Book *Book `json:"book"`
}

type BooksResponse struct {
	Book *[]Book `json:"books"`
}

type ReadinglistModel struct {
	Endpoint string
}
