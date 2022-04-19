package handler

import (
	"log"
	"net/http"
	"testing"

	"github.com/digyx/vorona/internal/sqlite"
)

var testChiRouter http.Handler

func TestMain(t *testing.M) {
	db, err := sqlite.New("../../../vorona.db")
	if err != nil {
		log.Fatal(err)
	}

	testChiRouter = New(db)

	if exitCode := t.Run(); exitCode != 0 {
		log.Fatal(exitCode)
	}
}
