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
	Books *[]Book `json:"books"`
}

type ReadinglistModel struct {
	Endpoint string
}

func (m *ReadinglistModel) GetAll() (*[]Book, error) {
	resp, err := http.Get(m.Endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var bookResp BooksResponse
	err = json.Unmarshal(data, &booksResp)
	if err != nil {
		return nil, err
	}

	return booksResp.Books, nil
}
