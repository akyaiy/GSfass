package config_test

import (
	"os"
	"testing"

	"github.com/akyaiy/GSfass/core/config"
)

type CfgConfig struct {
	Field1 string `mapstructure:"field1"`
	Field2 int    `mapstructure:"field2"`
	Field3 int    `mapstructure:"field3"`
}

func Test_configReadingUsingString(t *testing.T) {
	tests := []struct {
		name   string
		cfgStr string
		expect CfgConfig
		typ    string
		defs   map[string]string
	}{
		{
			name: "yaml_basic",
			cfgStr: `
field1: "testValue"
field2: 123
`,
			expect: CfgConfig{
				Field1: "testValue",
				Field2: 123,
			},
			typ: "yaml",
		},
		{
			name: "json_basic_with_defaults",
			cfgStr: `
{
	"field1": "jsonValue",
	"field2": 456
}
`,
			expect: CfgConfig{
				Field1: "jsonValue",
				Field2: 456,
				Field3: 789,
			},
			typ: "json",
			defs: map[string]string{
				"field3": "789",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &CfgConfig{}
			err := config.Read().Config().String(tt.cfgStr).SetType(tt.typ).
				SetBy(cfg).
				SetDefaults(tt.defs).
				End()

			if err != nil {
				t.Fatalf("Failed to read config from string: %v", err)
			}
			if *cfg != tt.expect {
				t.Errorf("Expected %+v, got %+v", tt.expect, *cfg)
			}
		})
	}
}


func Test_configReadingUsingFilepath(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expect   CfgConfig
		defs     map[string]string
	}{
		{
			name:     "yaml_file_basic",
			filePath: "testdata/config.yaml",
			expect: CfgConfig{
				Field1: "testValue",
				Field2: 123,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &CfgConfig{}
			err := config.Read().Config().FilePath(tt.filePath).
				SetBy(cfg).
				SetDefaults(tt.defs).
				End()

			if err != nil {
				t.Fatalf("Failed to read config from file path: %v", err)
			}
			if *cfg != tt.expect {
				t.Errorf("Expected %+v, got %+v", tt.expect, *cfg)
			}
		})
	}
}

func Test_configReadingUsingFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expect   CfgConfig
		typ      string
		defs     map[string]string
	}{
		{
			name:     "yaml_reader_basic",
			filePath: "testdata/config.yaml",
			expect: CfgConfig{
				Field1: "testValue",
				Field2: 123,
			},
			typ: "yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &CfgConfig{}

			file, err := os.Open(tt.filePath)
			if err != nil {
				t.Fatalf("Failed to open config file: %v", err)
			}
			defer file.Close()

			err = config.Read().Config().File(*file).SetType(tt.typ).
				SetBy(cfg).
				SetDefaults(tt.defs).
				End()

			if err != nil {
				t.Fatalf("Failed to read config from file: %v", err)
			}
			if *cfg != tt.expect {
				t.Errorf("Expected %+v, got %+v", tt.expect, *cfg)
			}
		})
	}
}