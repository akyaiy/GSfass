package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type EnvReader struct {
	by     any
	defs   map[string]string
	prefix string
}

func (r Readers) Environment() EnvReader {
	return EnvReader{}
}

func (er EnvReader) SetBy(sct any) EnvReader {
	er.by = sct
	return er
}

func (er EnvReader) SetDefaults(defs map[string]string) EnvReader {
	er.defs = defs
	return er
}

func (er EnvReader) SetEnvPrefix(prefix string) EnvReader {
	er.prefix = prefix
	return er
}

func (er EnvReader) End() error {
	if er.by == nil {
		return fmt.Errorf("no source provided for environment reading")
	}

	return er.read()
}

func (er EnvReader) read() error {
	v := viper.New()

	keys, err := extractMapstructureKeys(er.by)
	if err != nil {
		return err
	}

	for _, key := range keys {
		if er.defs != nil {
			if def, ok := er.defs[key]; ok {
				v.SetDefault(key, def)
			}
		}

		v.BindEnv(key)
	}

	if er.prefix != "" {
		v.SetEnvPrefix(er.prefix)
	}

	v.AutomaticEnv()

	if err := v.Unmarshal(er.by); err != nil {
		return fmt.Errorf("failed to unmarshal environment variables: %w", err)
	}

	return nil
}

func extractMapstructureKeys(obj any) ([]string, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("object must be struct or *struct")
	}

	var keys []string

	var walk func(t reflect.Type, prefix string)
	walk = func(t reflect.Type, prefix string) {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			tag := f.Tag.Get("mapstructure")
			if tag == "" {
				continue
			}

			full := tag
			if prefix != "" {
				full = prefix + "." + tag
			}

			if f.Type.Kind() == reflect.Struct {
				walk(f.Type, full)
			} else {
				keys = append(keys, full)
			}
		}
	}

	walk(t, "")

	return keys, nil
}
