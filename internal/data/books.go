package data

import (
	"database/sql"
	"errors"
	"time"
	"github.com/lib/pq"
)

type Book struct {
	ID int64			`json:"id"`
	CreatedAt time.Time `json:"-"`
	Title string 		`json:"title"`
	Published int 		`json:"published,omitempty"`
	Pages int 			`json:"pages,omitempty"`	
	Genres []string		`json:"genres,omitempty"`
	Rating float32		`json:"raring,omitempty"`
	Version int32		`json:"-"`
}

type MokModel struct {
	DB *sql.DB
}

func (b BookModel) Insert(book *Book) error {
	query := `
	INSERT INTO books (title, published, pages, genres, rating)
	VALUES ($1, $2, $3, $4, $5)
	RETURENING id, crated_at, version`

	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.genres), book.Rating}
	// return the auto generated system values to Go object
	return b.DB.QueryRow(query, args...).Scan(&book.ID, &book.CreatedAt, &book.Version)
}

func (b BookModel) Get(id int64) (*Book, error) {
	if id < 1 {
		return nil, errors.New("record not found")
	}

	query := `
	SELECT id, created_at, title, published, pages, genres, rating, version
	FROM books
	WHERE id = $1`

	var book Book

	err := b.DB.QueryRow(query, id).Scan(
		&book.ID,
		&book.CreatedAt,
		&book.Title,
		&book.Published,
		&book.Pages,
		pq.Array(&book.Genres),
		&book.Rating,
		&book.Version,
	)

	if err != nil {
		switch  {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &book, nil
}

func (b BookModel) Update(book *Book) error {
	query := `
	UPDATE books
	SET title = $1, published = $2, pages = $3, genres = $4, rating = $5, version = version + 1
	RETURNING version`

	args := []interface{}{book.Title, book.Published, book.Pages, pq.Array(book.genres), book.Rating, book.ID}
	return b.DB.QueryRow(query, args...).Scan(&book.Version)
}

func (b BookModel) Delete(id, int64) error {
	if id < 1 {
		return errors.New("record not found")
	}

	query := `
	DELETE FROM books
	WHERE id = $1`

	results, err := b.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected === 0 {
		return errors.New("record not found") 
	}
}