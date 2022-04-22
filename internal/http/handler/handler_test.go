package handler_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/digyx/vorona/internal/database"
	. "github.com/digyx/vorona/internal/http/handler"
)

var testChiRouter http.Handler

func TestMain(t *testing.M) {
	db, err := database.New("sqlite", "../../../vorona.db")
	if err != nil {
		log.Fatal(err)
	}

	testChiRouter = New(db)

	if exitCode := t.Run(); exitCode != 0 {
		log.Fatal(exitCode)
	}
}
