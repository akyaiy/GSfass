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
		name   string
		env    map[string]string
		defs   map[string]string
		prefix string
		expect EnvConfig
	}{
		{
			name: "basic environment reading with defaults",
			env: map[string]string{
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
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				t.Setenv(k, v)
			}

			cfg := &EnvConfig{}

			if err := config.Read().
				Environment().
				SetBy(cfg).
				SetDefaults(tt.defs).
				SetEnvPrefix(tt.prefix).
				End(); err != nil {
				t.Fatalf("failed to read environment variables: %v", err)
			}

			// Проверяем результат
			if *cfg != tt.expect {
				t.Errorf("expected %+v, got %+v", tt.expect, *cfg)
			}
		})
	}
}
