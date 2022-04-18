package sqlite

import (
	"reflect"
	"testing"

	"github.com/digyx/vorona/internal"
	"github.com/digyx/vorona/mock"
)

func TestAllBooks(t *testing.T) {
	expected := []internal.Book{
		mock.AzureWitch,
		mock.BloodOath,
	}

	db, err := New("../../../vorona.db")
	if err != nil {
		t.Error(err)
	}

	result, err := db.AllBooks()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nwant %v\ngot  %v", expected, result)
	}
}

func TestBook(t *testing.T) {
	expected := mock.AzureWitch

	db, err := New("../../../vorona.db")
	if err != nil {
		t.Error(err)
	}

	result, err := db.Book("AzureWitch")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nwant %v\ngot  %v", expected, result)
	}
}
