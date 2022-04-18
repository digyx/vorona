package handler

import (
	"reflect"
	"testing"

	"github.com/digyx/vorona/internal"
)

// These needs mocks
func TestBookToMustache(t *testing.T) {
	tests := []struct {
		name   string
		fields internal.Book
		want   map[string]interface{}
	}{
		{
			// Test that midnight UTC display as the correct day
			// Timezones are a pain
			name: "12:00:00AM UTC",
			fields: internal.Book{
				Slug:        "AzureWitch",
				Title:       "Death of the Azure Witch",
				Description: "This is a real book.",
				ReleaseTime: 10413792000,
			},
			want: map[string]interface{}{
				"slug":         "AzureWitch",
				"title":        "Death of the Azure Witch",
				"description":  "<p>This is a real book.</p>",
				"is_released":  false,
				"release_date": "January 01 2300",
			},
		},
		{
			// Same as above but the previous day
			name: "11:59:59PM UTC",
			fields: internal.Book{
				Slug:        "AzureWitch",
				Title:       "Death of the Azure Witch",
				Description: "This is a real book.",
				ReleaseTime: 10413791999,
			},
			want: map[string]interface{}{
				"slug":         "AzureWitch",
				"title":        "Death of the Azure Witch",
				"description":  "<p>This is a real book.</p>",
				"is_released":  false,
				"release_date": "December 31 2299",
			},
		},
		{
			// Test that books are marked release if it's passed now
			name: "Released",
			fields: internal.Book{
				Slug:        "AzureWitch",
				Title:       "Death of the Azure Witch",
				Description: "This is a real book.",
				ReleaseTime: 0,
			},
			want: map[string]interface{}{
				"slug":         "AzureWitch",
				"title":        "Death of the Azure Witch",
				"description":  "<p>This is a real book.</p>",
				"is_released":  true,
				"release_date": "January 01 1970",
			},
		},
		{
			// Ensure that markdown is properly rendered
			name: "Markdown",
			fields: internal.Book{
				Slug:        "AzureWitch",
				Title:       "Death of the Azure Witch",
				Description: "*This* is a **real** book.",
				ReleaseTime: 0,
			},
			want: map[string]interface{}{
				"slug":         "AzureWitch",
				"title":        "Death of the Azure Witch",
				"description":  "<p><em>This</em> is a <strong>real</strong> book.</p>",
				"is_released":  true,
				"release_date": "January 01 1970",
			},
		},
	}

	// Run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book := &internal.Book{
				Slug:        tt.fields.Slug,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				ReleaseTime: tt.fields.ReleaseTime,
			}
			if got := BookToMustache(book); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookToMustache() = %v, want %v", got, tt.want)
			}
		})
	}
}
