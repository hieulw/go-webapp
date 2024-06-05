package main

import (
	"log"
	"ticket/internal/server"
)

func main() {
	server := server.NewServer()

	if err := server.Start(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
