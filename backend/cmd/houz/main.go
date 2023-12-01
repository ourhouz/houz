package main

import (
	"github.com/ourhouz/houz/internal/config"
	"github.com/ourhouz/houz/internal/db"
	"github.com/ourhouz/houz/internal/server"
)

func main() {
	config.Load()

	db.Connect()
	db.Init()

	server.Init()
}
