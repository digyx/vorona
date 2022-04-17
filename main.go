package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/digyx/vorona/internal/db/sqlite"
	"github.com/digyx/vorona/internal/http/handler"
	"github.com/digyx/vorona/internal/http/server"
	"github.com/digyx/vorona/mock"
)

func main() {
	app := &cli.App{
		Name:  "vorona",
		Usage: "Webserver for vorona.gg",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "sqlite",
				Usage: "Path to the SQLite database",
			},
			&cli.StringFlag{
				Name:  "bind",
				Value: "0.0.0.0:8080",
				Usage: "IP Address the server will attempt to bind to",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "start",
				Usage:  "Start the webserver",
				Action: run,
			},
			{
				Name:  "dev-db",
				Usage: "Rebuild the SQLlite DB for development",
				Action: func(c *cli.Context) error {
					path := c.String("sqlite")
					if path == "" {
						return fmt.Errorf("error: sqlite flag is required")
					}

					return mock.RebuildDevelopmentDatabase(path)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	sqlitePath := c.String("sqlite")
	if sqlitePath == "" {
		return fmt.Errorf("error: sqlite flag is required")
	}

	database, err := sqlite.New(sqlitePath)
	if err != nil {
		log.Print("error: could not connect to database")
		return err
	}

	// Initialize Handler
	httpHandler := handler.New(database)

	// Setup and run webserver
	server, cancel := server.New(c.String("bind"), httpHandler)
	defer cancel()

	// This is blocking
	server.ListenAndServe()
	return nil
}
