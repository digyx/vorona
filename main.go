package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/digyx/vorona/internal/database"
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
				Name:  "postgres",
				Usage: "Connection URI for a Postgres Database (priority over SQLite)",
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
			// Used during development for a consistent environment
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

// Actually start the webserver with a SQLite database
func run(c *cli.Context) error {
	var driver string
	var dbPath string

	// Determine which database to connect to
	if postgresURI := c.String("postgres"); postgresURI != "" {
		driver, dbPath = "pgx", postgresURI
	} else if sqlitePath := c.String("sqlite"); sqlitePath != "" {
		driver, dbPath = "sqlite", sqlitePath
	} else {
		log.Fatal("error: ether the --sqlite or --postgres flag must be passed")
	}

	// Terminate on error
	db, err := database.New(driver, dbPath)
	if err != nil {
		log.Print("error: could not connect to database")
		return err
	}

	// Initialize Handler
	httpHandler := handler.New(db)

	// Setup the webserver
	server := server.New(c.String("bind"), httpHandler)

	// This is blocking
	server.ListenAndServe()
	return nil
}
