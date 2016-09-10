package controllers

import (
	"bookstore/models"
	//	"bookstore/server"
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

	host := ""
	port := "3000"

	printHost := host
	if printHost == "" {
		printHost = "localhost"
	}
	ctlr.server.Logger.Print("Listening on " + printHost + ":" + port)

	http.ListenAndServe(host+":"+port, nil)
}

func NewController(svr *server.Server, ds models.Datastore) Controller {
	return &controller{server: svr, datastore: ds}
}
