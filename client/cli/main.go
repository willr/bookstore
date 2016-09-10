package main

import (
	//	"bookstore/config"
	//	"bookstore/models"
	// 	"bookstore/client"
	"log"
	"os"
)

var (
// 	client  *client.Client
)

func main() {

	// 	conf := config.Parse()
	log := buildLogger()

	log.Println("Nothing to do...")
	// log.Fatal(os.Stdout, "Nothing to do...")
}

func buildLogger() *log.Logger {

	l := log.New(os.Stdout, "bookstore-client: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	return l
}
