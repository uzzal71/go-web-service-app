package data

import (
	"time"
)

type Book struct {
	ID int64	`json:"id"`
	CreatedAt time.Time
	Title string 
	Published int 
	Pages int 
	Genres []string
	Rating float32
	Version int32
}