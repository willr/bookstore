package models

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type Datastore interface {
	AllBooks() ([]*Book, error)
	BookByIsbn(isbn string) (*Book, error)
	CreateBook(isbn string, title string, author string, price float64) (int64, error)
}

type Book struct {
	Isbn   string
	Title  sql.NullString
	Author sql.NullString
	Price  sql.NullFloat64
}

func (db *DB) AllBooks() ([]*Book, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func (db *DB) BookByIsbn(isbn string) (*Book, error) {

	if isbn == "" {
		return nil, errors.New("invalid isbn param")
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := new(Book)
	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err == sql.ErrNoRows {
		return nil, errors.New("isdn not found")
	} else if err != nil {
		return nil, err
	}

	return bk, nil
}

func (db *DB) CreateBook(isbn string, title string, author string, price float64) (int64, error) {

	result, err := db.Exec("INSERT INTO books VALUES($1, $2, $3, $4)", isbn, title, author, price)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
