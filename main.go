package main

import (
	"bookstore/config"
	"bookstore/controllers"
	"bookstore/models"
	"bookstore/server"
	"log"
	"os"
)

var (
	srv  *server.Server
	ctlr controllers.Controller
)

func main() {

	conf := config.Parse()
	log := buildLogger()
	srv = &server.Server{log, conf}

	db := models.SetupDB(conf)

	ctlr = controllers.NewController(srv, db)
	ctlr.SetupControllers()

	ctlr.RunServer()
}

func buildLogger() *log.Logger {

	l := log.New(os.Stdout, "bookstore: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	return l
}
