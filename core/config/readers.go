// The config package is part of the core module.
// It provides functions for reading environment variables,
// configuration files, and string parameters.

package config

import "os"

type Readers struct{}

func Read() Readers {
	return Readers{}
}

type ReaderContract interface {
	Environment() EnvReaderContract
	Config() CfgReaderContract
	// CmdLine() CmdLineReaderContract
}

type EnvReaderContract interface {
	SetBy(sct any) EnvReaderContract
	SetDefaults(defs map[string]string) EnvReaderContract
	SetEnvPrefix(prefix string) EnvReader
	End() error
}

type CfgReaderContract interface {
	SetType(typ string) CfgReaderContract
	File(file os.File) CfgReaderContract
	FilePath(path string) CfgReaderContract
	String(str string) CfgReaderContract
	SetBy(sct any) CfgReaderContract
	SetDefaults(defs map[string]any) CfgReaderContract
	End() error
}

// type CmdLineReaderContract interface {
// 	SetBy(sct any) CmdLineReaderContract
// 	SetQuery(query []string) CmdLineReaderContract
// 	SetOpts(opts CmdLineReaderOptions) CmdLineReader
// 	End() error
// }
