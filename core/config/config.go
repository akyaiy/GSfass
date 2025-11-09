package config

import (
	"fmt"
	"os"
)

const (
	UsingFilePath = iota
	UsingFile
	UsingString
)

type ConfigReader struct {
	by       any
	defs     map[string]any
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

func (cr ConfigReader) SetFile(file os.File) ConfigReader {
	cr.file = file
	cr.using = UsingFile
	return cr
}

func (cr ConfigReader) SetFilePath(path string) ConfigReader {
	cr.filePath = path
	cr.using = UsingFilePath
	return cr
}

func (cr ConfigReader) SetString(str string) ConfigReader {
	cr.str = str
	cr.using = UsingString
	return cr
}

func (cr ConfigReader) SetDefaults(defs map[string]any) ConfigReader {
	cr.defs = defs
	return cr
}

func (cr ConfigReader) End() error {
	if cr.by == nil {
		return fmt.Errorf("no source provided for config reading")
	}
	// TODO: Implement config reading logic here
	return nil
}
