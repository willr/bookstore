package main

import (
	"bookstore/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Env struct {
	bookDatastore models.BookDatastore
}

func main() {

	db, err := models.NewDB("postgres://user:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/books", env.booksIndex)
	http.HandleFunc("/books/show", env.booksShow)
	http.HandleFunc("/books/create", env.booksCreate)
	http.ListenAndServe(":3000", nil)
}

func getBookValues(w http.ResponseWriter, bks []*models.Book) {

	// printSlice("ss", bks)
	for _, bk := range bks {
		var price string

		if bk.Price.Valid {
			price = fmt.Sprintf("%.2f", bk.Price.Float64)
		} else {
			price = "PRICE NOT SET"
		}
		fmt.Fprintf(w, "%s, %s, %s, %s\n", bk.Isbn, bk.Title.String, bk.Author.String, price)
	}

}

func printSlice(s string, x []*models.Book) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	bks, _ := env.bookDatastore.AllBooks()
	getBookValues(w, bks)
}

func (env *Env) booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	bk, _ := env.bookDatastore.BookByIsbn(isbn)

	bks := make([]*models.Book, 0)
	bks = append(bks, bk)
	getBookValues(w, bks)

}

func (env *Env) showBookByIsbn(w http.ResponseWriter, r *http.Request, isbn string) {

	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	bk, _ := env.bookDatastore.BookByIsbn(isbn)

	bks := make([]*models.Book, 0)
	bks = append(bks, bk)
	getBookValues(w, bks)

}

func (env *Env) booksCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	title := r.FormValue("title")
	author := r.FormValue("author")
	if isbn == "" || title == "" || author == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rowsAffected, err := env.bookDatastore.CreateBook(isbn, title, author, price)

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
	env.showBookByIsbn(w, r, isbn)

}
