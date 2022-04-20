package mock

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/digyx/vorona/internal"
)

var (
	AzureWitch internal.Book = internal.Book{
		Slug:        "AzureWitch",
		Title:       "Death of the Azure Witch",
		Description: "This is a real book.",
		ReleaseTime: 0,
	}

	BloodOath internal.Book = internal.Book{
		Slug:        "BloodOath",
		Title:       "Blood Oath",
		Description: "Sometimes, someone needs to die.",
		ReleaseTime: 1644796800,
	}

	MidnightRelease internal.Book = internal.Book{
		Slug:        "MidnightRain",
		Title:       "Midnight Rain",
		Description: "Amaya Kuroshi",
		ReleaseTime: 10413792000,
	}

	EleventhHour internal.Book = internal.Book{
		Slug:        "EleventhHour",
		Title:       "Dri Daltan",
		Description: "Time stood still.",
		ReleaseTime: 10413791999,
	}

	MarkOfInsanity internal.Book = internal.Book{
		Slug:        "MarkOfInsanity",
		Title:       "Mark of Insanity",
		Description: "**This** is *the* moment.",
		ReleaseTime: 1,
	}

	AllBookMocks []internal.Book = []internal.Book{
		MidnightRelease,
		EleventhHour,
		BloodOath,
		MarkOfInsanity,
		AzureWitch,
	}
)

func RebuildDevelopmentDatabase(path string) error {
	// First arg is the path to the DB
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil
	}

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	tx.Exec("DROP TABLE books")

	createTableStmt := `
	CREATE TABLE books (
		book_id      INTEGER PRIMARY KEY,
		slug         TEXT    NOT NULL UNIQUE,
		title        TEXT    NOT NULL,
		description  TEXT    NOT NULL,
		release_time INTEGER NOT NULL
	)`

	insertStmt, err := tx.Prepare("INSERT INTO books (slug, title, description, release_time) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	tx.Exec(createTableStmt)

	for _, book := range AllBookMocks {
		insertStmt.Exec(book.Slug, book.Title, book.Description, book.ReleaseTime)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
