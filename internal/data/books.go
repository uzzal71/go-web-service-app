package data

import (
	"time"
)

type Book struct {
	ID int64			`json:"id"`
	CreatedAt time.Time `json:"-"`
	Title string 		`json:"title"`
	Published int 		`json:"published,omitempty"`
	Pages int 			`json:"oages,omitempty"`	
	Genres []string		`json:"genres,omitempty"`
	Rating float32		`json:"raring,omitempty"`
	Version int32		
}