package handler

import (
	"reflect"
	"testing"

	"github.com/digyx/vorona/internal"
)

func TestBookToMustache(t *testing.T) {
	type fields struct {
		Slug        string
		Title       string
		Description string
		ReleaseTime int64
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "12:00:00AM UTC",
			fields: fields{
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
			name: "11:59:59PM UTC",
			fields: fields{
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
			name: "Released",
			fields: fields{
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
			name: "Markdown",
			fields: fields{
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
