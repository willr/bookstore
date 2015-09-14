package controllers

import (
	"bookstore/models"
	"fmt"
	"net/http"
	"strconv"
)

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

func (ctlr *controller) BookIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	fmt.Println("function pointers: ", ctlr)
	bks, _ := ctlr.datastore.AllBooks()
	getBookValues(w, bks)
}

func (ctlr *controller) BookShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	bk, _ := ctlr.datastore.BookByIsbn(isbn)

	bks := make([]*models.Book, 0)
	bks = append(bks, bk)
	getBookValues(w, bks)

}

func (ctlr *controller) ShowBookByIsbn(w http.ResponseWriter, r *http.Request, isbn string) {

	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	bk, _ := ctlr.datastore.BookByIsbn(isbn)

	bks := make([]*models.Book, 0)
	bks = append(bks, bk)
	getBookValues(w, bks)

}

func (ctlr *controller) BookCreate(w http.ResponseWriter, r *http.Request) {
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

	rowsAffected, err := ctlr.datastore.CreateBook(isbn, title, author, price)

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
	ctlr.ShowBookByIsbn(w, r, isbn)

}
