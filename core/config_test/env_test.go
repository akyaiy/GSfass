package config_test

import (
	"testing"

	"github.com/akyaiy/GSfass/core/config"
)

type EnvConfig struct {
	Var1 string `mapstructure:"VAR1"`
	Var2 string `mapstructure:"VAR2"`
	Var3 string `mapstructure:"VAR3"`
}

func Test_envReading(t *testing.T) {
	tests := []struct {
		obj    any
		eviron map[string]string
		defs   map[string]string
		prefix string
		expect EnvConfig
	}{
		{
			obj: &EnvConfig{},
			eviron: map[string]string{
				"VAR1": "value1",
				"VAR2": "value2",
			},
			defs: map[string]string{
				"VAR3": "default3",
			},
			prefix: "",
			expect: EnvConfig{
				Var1: "value1",
				Var2: "value2",
				Var3: "default3",
			},
		},
	}

	for _, tt := range tests {
		for k, v := range tt.eviron {
			t.Setenv(k, v)
		}

		cfg := &EnvConfig{}

		err := config.Read().Environment().SetBy(cfg).
			SetDefaults(tt.defs).
			SetEnvPrefix(tt.prefix).
			End()

		if err != nil {
			t.Fatalf("Failed to read environment variables: %v", err)
		}

		if *cfg != tt.expect {
			t.Errorf("Expected %+v, got %+v", tt.expect, *cfg)
		}
	}
}
