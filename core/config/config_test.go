package config

import (
	"testing"
)

func Test_extractMapstructureKeys(t *testing.T) {
	tests := []struct {
		obj    any
		expect []string
	}{
		{
			obj: struct {
				Field1 string `mapstructure:"field_1"`
				Field2 int    `mapstructure:"field_2"`
			}{},
			expect: []string{"field_1", "field_2"},
		},
		{
			obj: &struct {
				Nested struct {
					SubField string `mapstructure:"sub_field"`
				} `mapstructure:"nested"`
			}{},
			expect: []string{"nested.sub_field"},
		},
	}

	for _, tt := range tests {
		keys, err := extractMapstructureKeys(tt.obj)
		if err != nil {
			t.Errorf("extractMapstructureKeys(%v) returned error: %v", tt.obj, err)
			continue
		}
		if len(keys) != len(tt.expect) {
			t.Errorf("extractMapstructureKeys(%v) = %v, want %v", tt.obj, keys, tt.expect)
			continue
		}
		for i, key := range keys {
			if key != tt.expect[i] {
				t.Errorf("extractMapstructureKeys(%v) = %v, want %v", tt.obj, keys, tt.expect)
				break
			}
		}
	}

	//extractMapstructureKeys(obj any) ([]string, error)
}
