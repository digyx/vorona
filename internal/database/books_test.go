package database_test

import (
	"testing"

	. "github.com/digyx/vorona/internal/database"
	"github.com/digyx/vorona/mock"
	"github.com/stretchr/testify/require"
)

// Does whatcha think; ensures the AllBooks and Book functions work as expected
func TestAllBooks(t *testing.T) {
	expected := mock.AllBookMocks

	db, err := New("sqlite", "../../vorona.db")
	if err != nil {
		t.Error(err)
	}

	result, err := db.GetAllBooks()
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, expected, result)
}

// Ensure internal.Book decode works
func TestBook(t *testing.T) {
	expected := mock.AzureWitch

	db, err := New("sqlite", "../../vorona.db")
	if err != nil {
		t.Error(err)
	}

	result, err := db.GetBook("AzureWitch")
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, expected, result)
}
