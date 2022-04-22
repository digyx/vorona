package handler_test

import (
	"reflect"
	"testing"

	"github.com/digyx/vorona/internal"
	. "github.com/digyx/vorona/internal/http/handler"
	"github.com/digyx/vorona/mock"
)

// Ensure the book -> template context translation is correct
func TestBookToMustache(t *testing.T) {
	tests := []struct {
		name   string
		fields internal.Book
		want   map[string]interface{}
	}{
		{
			// Test that midnight UTC display as the correct day
			// Timezones are a pain
			name:   "12:00:00AM UTC",
			fields: mock.MidnightRelease,
			want: map[string]interface{}{
				"slug":         "MidnightRain",
				"title":        "Midnight Rain",
				"description":  "<p>Amaya Kuroshi</p>",
				"is_released":  false,
				"release_date": "January 01 2300",
			},
		},
		{
			// Same as above but the previous day
			name:   "11:59:59PM UTC",
			fields: mock.EleventhHour,
			want: map[string]interface{}{
				"slug":         "EleventhHour",
				"title":        "Dri Daltan",
				"description":  "<p>Time stood still.</p>",
				"is_released":  false,
				"release_date": "December 31 2299",
			},
		},
		{
			// Test that books are marked release if it's passed now
			name:   "Released",
			fields: mock.AzureWitch,
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
			name:   "Markdown",
			fields: mock.MarkOfInsanity,
			want: map[string]interface{}{
				"slug":         "MarkOfInsanity",
				"title":        "Mark of Insanity",
				"description":  "<p><strong>This</strong> is <em>the</em> moment.</p>",
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
