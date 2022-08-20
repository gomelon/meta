package meta

import (
	"reflect"
	"testing"
)

func TestMeta_MapStruct(t *testing.T) {
	type fields struct {
		name       string
		properties map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		input  any
		wanted any
	}{
		{
			name: "should map struct",
			fields: fields{
				name: "Test",
				properties: map[string]string{
					"StrField":     "string",
					"IntField":     "1",
					"Int64Field":   "1",
					"Uint64Field":  "1",
					"Int32Field":   "1",
					"Uint32Field":  "1",
					"BoolField1":   "true",
					"BoolField2":   "BoolField2",
					"Float32Field": "1",
					"Float64Field": "1",
				},
			},
			input: &MetaStruct{},
			wanted: &MetaStruct{
				StrField:     "string",
				IntField:     1,
				Int64Field:   1,
				Uint64Field:  1,
				Int32Field:   1,
				Uint32Field:  1,
				BoolField1:   true,
				BoolField2:   true,
				Float32Field: 1,
				Float64Field: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Meta{
				name:       tt.fields.name,
				properties: tt.fields.properties,
			}
			m.MapStruct(tt.input)

			if !reflect.DeepEqual(tt.input, tt.wanted) {
				t.Errorf("MapStruct() = %v, want %v", tt.input, tt.wanted)
			}
		})
	}
}

type MetaStruct struct {
	StrField     string
	IntField     int
	Int64Field   int64
	Uint64Field  uint64
	Int32Field   int32
	Uint32Field  uint32
	BoolField1   bool
	BoolField2   bool
	Float32Field float32
	Float64Field float64
}
