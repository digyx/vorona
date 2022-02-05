package books

import (
	"reflect"
	"testing"
)

func TestBook_ToMap(t *testing.T) {
	type fields struct {
		ID          string
		Title       string
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
				ID:          "AzureWitch",
				Title:       "Death of the Azure Witch",
				ReleaseTime: 1646006400,
			},
			want: map[string]interface{}{
				"book.id":           "AzureWitch",
				"book.title":        "Death of the Azure Witch",
				"book.is_released":  false,
				"book.release_date": "February 28 2022",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			book := &Book{
				ID:          tt.fields.ID,
				Title:       tt.fields.Title,
				ReleaseTime: tt.fields.ReleaseTime,
			}
			if got := book.ToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Book.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
