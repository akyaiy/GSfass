package config

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	UsingFilePath = iota
	UsingFile
	UsingString
)

type ConfigReader struct {
	by       any
	defs     map[string]string
	typ      string
	using    int8
	file     os.File
	filePath string
	str      string
}

func (r Readers) Config() ConfigReader {
	return ConfigReader{}
}

func (cr ConfigReader) SetBy(sct any) ConfigReader {
	cr.by = sct
	return cr
}

func (cr ConfigReader) SetType(typ string) ConfigReader {
	cr.typ = typ
	return cr
}

func (cr ConfigReader) File(file os.File) ConfigReader {
	cr.file = file
	cr.using = UsingFile
	return cr
}

func (cr ConfigReader) FilePath(path string) ConfigReader {
	cr.filePath = path
	cr.using = UsingFilePath
	return cr
}

func (cr ConfigReader) String(str string) ConfigReader {
	cr.str = str
	cr.using = UsingString
	return cr
}

func (cr ConfigReader) SetDefaults(defs map[string]string) ConfigReader {
	cr.defs = defs
	return cr
}

func (cr ConfigReader) End() error {
	if cr.by == nil {
		return fmt.Errorf("no source provided for config reading")
	}
	if cr.using == UsingFilePath && cr.filePath == "" {
		return fmt.Errorf("file path is empty")
	}
	if cr.using == UsingFile && cr.typ == "" {
		return fmt.Errorf("config type is not specified")
	} else if cr.using == UsingString && cr.typ == "" {
		return fmt.Errorf("config type is not specified")
	}
	return cr.read()
}

func (cr ConfigReader) read() error {
	v := viper.New()

	for k, def := range cr.defs {
		v.SetDefault(k, def)
	}

	switch cr.using {
	case UsingFilePath:
		v.SetConfigFile(cr.filePath)
		if err := v.ReadInConfig(); err != nil {
			return fmt.Errorf("cannot read config from file path: %w", err)
		}
	case UsingFile:
		v.SetConfigType(cr.typ)
		if cr.file.Fd() == 0 && cr.file.Fd() <= 2 {
			return fmt.Errorf("provided file is not valid")
		}

		b, err := os.ReadFile(cr.file.Name())
		if err != nil {
			return fmt.Errorf("cannot read config from file: %w", err)
		}

		if err := v.ReadConfig(bytes.NewReader(b)); err != nil {
			return fmt.Errorf("cannot read config from file: %w", err)
		}
	case UsingString:
		v.SetConfigType(cr.typ)
		if err := v.ReadConfig(strings.NewReader(cr.str)); err != nil {
			return fmt.Errorf("cannot read config from string: %w", err)
		}
	default:
		return fmt.Errorf("no valid source provided for config reading")
	}

	if err := v.Unmarshal(cr.by); err != nil {
		return fmt.Errorf("cannot unmarshal config into struct: %w", err)
	}
	return nil
}
