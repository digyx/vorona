package handler_test

import (
	"testing"

	"github.com/digyx/vorona/internal"
	. "github.com/digyx/vorona/internal/http/handler"
	"github.com/digyx/vorona/mock"
	"github.com/stretchr/testify/require"
)

// Ensure the book -> template context translation is correct
func TestBookToMustache(t *testing.T) {
	tests := []struct {
		name     string
		book     internal.Book
		expected map[string]interface{}
	}{
		{
			// Test that midnight UTC display as the correct day
			// Timezones are a pain
			name: "12:00:00AM UTC",
			book: mock.MidnightRelease,
			expected: map[string]interface{}{
				"slug":         "MidnightRain",
				"title":        "Midnight Rain",
				"description":  "<p>Amaya Kuroshi</p>",
				"is_released":  false,
				"release_date": "January 01 2300",
				"image_url":    "https://imagedelivery.net/15nTRPX67jTK1T3emsbaRA/1a620cdf-5baa-4f67-5341-f90fb19e7f00/public",
			},
		},
		{
			// Same as above but the previous day
			name: "11:59:59PM UTC",
			book: mock.EleventhHour,
			expected: map[string]interface{}{
				"slug":         "EleventhHour",
				"title":        "Dri Daltan",
				"description":  "<p>Time stood still.</p>",
				"is_released":  false,
				"release_date": "December 31 2299",
				"image_url":    "https://imagedelivery.net/15nTRPX67jTK1T3emsbaRA/1a620cdf-5baa-4f67-5341-f90fb19e7f00/public",
			},
		},
		{
			// Test that books are marked release if it's passed now
			name: "Released",
			book: mock.AzureWitch,
			expected: map[string]interface{}{
				"slug":         "AzureWitch",
				"title":        "Death of the Azure Witch",
				"description":  "<p>This is a real book.</p>",
				"is_released":  true,
				"release_date": "January 01 1970",
				"image_url":    "https://imagedelivery.net/15nTRPX67jTK1T3emsbaRA/1a620cdf-5baa-4f67-5341-f90fb19e7f00/public",
			},
		},
		{
			// Ensure that markdown is properly rendered
			name: "Markdown",
			book: mock.MarkOfInsanity,
			expected: map[string]interface{}{
				"slug":         "MarkOfInsanity",
				"title":        "Mark of Insanity",
				"description":  "<p><strong>This</strong> is <em>the</em> moment.</p>",
				"is_released":  true,
				"release_date": "January 01 1970",
				"image_url":    "https://imagedelivery.net/15nTRPX67jTK1T3emsbaRA/1a620cdf-5baa-4f67-5341-f90fb19e7f00/public",
			},
		},
	}

	// Run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BookToMustache(&tt.book)
			require.Equal(t, tt.expected, result)
		})
	}
}
