package sqlite

import (
	"reflect"
	"testing"

	"github.com/digyx/vorona/internal"
)

func TestAllBooks(t *testing.T) {
	expected := []internal.Book{
		{
			Slug:        "AzureWitch",
			Title:       "Death of the Azure Witch",
			Description: "This is a real book.",
			ReleaseTime: 1646006400,
		},
		{
			Slug:        "BloodOath",
			Title:       "Blood Oath",
			Description: "Sometimes, someone needs to die.",
			ReleaseTime: 1644796800,
		},
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
	expected := internal.Book{
		Slug:        "AzureWitch",
		Title:       "Death of the Azure Witch",
		Description: "This is a real book.",
		ReleaseTime: 1646006400,
	}

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
