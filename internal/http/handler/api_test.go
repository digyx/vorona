package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/digyx/vorona/internal/db/sqlite"
	"github.com/digyx/vorona/mock"
)

func executeTest(path string, expected string) error {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	db, err := sqlite.New("../../../vorona.db")
	if err != nil {
		return err
	}

	router := New(db)
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		return fmt.Errorf("expected status 200, got %d", status)
	}

	// Strip whitespace because Chi adds in a \n at the end of the result
	if body := recorder.Body.String(); strings.TrimSpace(body) != expected {
		return fmt.Errorf("incorrect body\nwant %s\ngot  %s", expected, body)
	}

	return nil
}

func TestApiBook(t *testing.T) {
	expected, _ := json.Marshal(mock.AzureWitch)
	err := executeTest("/api/book/AzureWitch", string(expected))

	if err != nil {
		t.Error(err)
	}
}
