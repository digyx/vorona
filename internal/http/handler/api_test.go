package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/digyx/vorona/mock"
	"github.com/stretchr/testify/require"
)

// All of these tests are basically the same, so I abstracted them
func executeTest(t *testing.T, path string, expected string) {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", path, nil)
	require.NoError(t, err)

	// Serve the HTTP request
	testChiRouter.ServeHTTP(recorder, req)

	status := recorder.Code
	require.Equal(t, http.StatusOK, status)

	// Strip whitespace because Chi adds in a \n at the end of the result
	body := strings.TrimSpace(recorder.Body.String())
	require.Equal(t, expected, body)
}

func TestApiBook(t *testing.T) {
	expected, _ := json.Marshal(mock.AzureWitch)
	executeTest(t, "/api/book/AzureWitch", string(expected))
}

func TestApiBooks(t *testing.T) {
	expected, _ := json.Marshal(mock.AllBookMocks)
	executeTest(t, "/api/book", string(expected))
}
