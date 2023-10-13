package data

import (
	"time"
)

type Book struct {
	ID int64	`json:"id"`
	CreatedAt time.Time `json:"-"`
	Title string 	`json:"title"`
	Published int 	`json:"published,omitempty"`
	Pages int 			`json:"published,omitempty"`	
	Genres []string		`json:"published,omitempty"`
	Rating float32		`json:"published,omitempty"`
	Version int32		`json:"published,omitempty"`
}