package controllers

import (
	"bookstore/models"
	"bookstore/server"
	"net/http"
)

type Controller interface {
	BookIndex(w http.ResponseWriter, r *http.Request)
	BookShow(w http.ResponseWriter, r *http.Request)
	BookCreate(w http.ResponseWriter, r *http.Request)

	SetupControllers()
	RunServer()
}

type controller struct {
	server    *server.Server
	datastore models.Datastore
}

func (ctlr controller) SetupControllers() {

	http.HandleFunc("/books", ctlr.BookIndex)
	http.HandleFunc("/books/show", ctlr.BookShow)
	http.HandleFunc("/books/create", ctlr.BookCreate)

	return
}

func (ctlr controller) RunServer() {

	http.ListenAndServe(":3000", nil)
}

func NewController(svr *server.Server, ds models.Datastore) Controller {
	return &controller{server: svr, datastore: ds}
}
