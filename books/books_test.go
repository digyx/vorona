package books

import (
	"reflect"
	"testing"
)

func TestBook_ToMustache(t *testing.T) {
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
			name: "basic",
			fields: fields{
				Slug:        "AzureWitch",
				Title:       "Death of the Azure Witch",
				Description: "This is a real book.",
				ReleaseTime: 1646006400,
			},
			want: map[string]interface{}{
				"slug":         "AzureWitch",
				"title":        "Death of the Azure Witch",
				"description":  "This is a real book.",
				"is_released":  false,
				"release_date": "February 28 2022",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book := &Book{
				Slug:        tt.fields.Slug,
				Title:       tt.fields.Title,
				ReleaseTime: tt.fields.ReleaseTime,
			}
			if got := book.ToMustache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Book.ToMustache() = %v, want %v", got, tt.want)
			}
		})
	}
}
