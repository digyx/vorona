package database_test

import (
	"reflect"
	"testing"

	. "github.com/digyx/vorona/internal/database"
	"github.com/digyx/vorona/mock"
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

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nwant %v\ngot  %v", expected, result)
	}
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

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nwant %v\ngot  %v", expected, result)
	}
}
