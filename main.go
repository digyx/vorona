package main

import (
	"fmt"

	"github.com/digyx/vorona/internal/db/sqlite"
	"github.com/digyx/vorona/internal/http/handler"
	"github.com/digyx/vorona/internal/http/server"
)

func main() {
	var err error
	database, err := sqlite.New("vorona.db")
	if err != nil {
		fmt.Println("error: could not connect to database")
		fmt.Println(err)
		return
	}

	// Initialize Handler
	httpHandler := handler.New(database)

	// Setup and run webserver
	server, cancel := server.New("0.0.0.0:8080", httpHandler)
	defer cancel()

	// This is blocking
	server.ListenAndServe()
}
