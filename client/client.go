package client

import (
	"bookstore/config"
	"log"
)

type Client struct {
	Logger *log.Logger
	Config *config.RuntimeConfig
}
