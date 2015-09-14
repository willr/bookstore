package server

import (
	"bookstore/config"
	"log"
)

type Server struct {
	Logger *log.Logger
	Config *config.RuntimeConfig
}
