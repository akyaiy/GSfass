package config

import (
	"fmt"
	"os"
)

type CmdLineReader struct {
	by    any
	query []string
}

func (r Readers) CmdLine() CmdLineReader {
	return CmdLineReader{}
}

func (clr CmdLineReader) SetBy(sct any) CmdLineReader {
	clr.by = sct
	return clr
}

func (clr CmdLineReader) SetQuery(query []string) CmdLineReader {
	clr.query = query
	return clr
}

func (clr CmdLineReader) End() error {
	if clr.by == nil {
		return fmt.Errorf("no source provided for command line reading")
	}
	if clr.query == nil {
		clr.query = os.Args
	}
	return clr.read()
}

func (clr CmdLineReader) read() error {
	return nil
}
