package main

import (
	"fmt"

	"github.com/digyx/vorona/internal/db"
	"github.com/digyx/vorona/internal/db/sqlite"
	"github.com/digyx/vorona/internal/http/server"
)

var database db.Database

func main() {
	var err error
	database, err = sqlite.New()
	if err != nil {
		fmt.Println("error: could not connect to database")
		fmt.Println(err)
		return
	}

	// Setup and run webserver
	server, cancel := server.New("0.0.0.0:8080", service())
	defer cancel()

	// This is blocking
	server.ListenAndServe()
}
