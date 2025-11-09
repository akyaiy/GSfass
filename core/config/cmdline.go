package config

// TODO: Implement command line reader

// import (
// 	"fmt"
// 	"os"
// 	"reflect"
// )

// const (
// 	LineUnitTypeFlag = iota
// 	LineUnitTypeParam
// 	LineUnitTypeCommand
// )

// type LineUnit struct {
// 	Must  bool
// 	Type  int
// 	Value any
// }

// type LineUnitMap map[string]LineUnit

// type CmdLineReaderOptions struct {
// 	IgnoreUnknown bool
// 	IgnoreCollisions bool
// 	UnitMap       *LineUnitMap
// }

// type CmdLineReader struct {
// 	by    any
// 	query []string
// 	units LineUnitMap
// 	opts  *CmdLineReaderOptions
// }

// func (r Readers) CmdLine() CmdLineReader {
// 	return CmdLineReader{}
// }

// func (clr CmdLineReader) SetBy(sct any) CmdLineReader {
// 	clr.by = sct
// 	return clr
// }

// func (clr CmdLineReader) SetQuery(query []string) CmdLineReader {
// 	clr.query = query
// 	return clr
// }

// func (clr CmdLineReader) SetOpts(opts CmdLineReaderOptions) CmdLineReader {
// 	clr.opts = &opts
// 	return clr
// }

// func (clr CmdLineReader) End() error {
// 	if clr.by == nil {
// 		return fmt.Errorf("no source provided for command line reading")
// 	}
// 	if clr.query == nil {
// 		clr.query = os.Args
// 	}
// 	if clr.opts == nil {
// 		clr.opts = &CmdLineReaderOptions{}
// 	} else if clr.opts.UnitMap != nil {
// 		clr.units = *clr.opts.UnitMap
// 	} else if clr.units == nil {
// 		clr.units = make(LineUnitMap)
// 	}
// 	return clr.read()
// }

// func (clr CmdLineReader) read() error {
// 	// units := make(LineUnitMap)
// 	// units["run"] = LineUnit{Type: LineUnitTypeCommand, Value: LineUnitMap{
// 	// 	"port": LineUnit{Type: LineUnitTypeParam, Value: nil},
// 	// 	"host": LineUnit{Type: LineUnitTypeParam, Value: nil},
// 	// }}

// 	return nil
// }

// func (clr CmdLineReader) unmarshal(target any, unitmap LineUnitMap) error {
// 	v := reflect.ValueOf(target)
// 	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
// 		return fmt.Errorf("cmdline unmarshal target must be a pointer to a struct")
// 	}
// 	return clr.decode(v.Elem(), unitmap)
// }

// // func (clr CmdLineReader) decode(v reflect.Value, unitmap LineUnitMap) error {
// // 	if v.Kind() != reflect.Struct {
// // 		return fmt.Errorf("expected struct, got %s", v.Kind())
// // 	}

// // 	t := v.Type()
// // 	for i := 0; i < t.NumField(); i++ {
// // 		field := t.Field(i)
// // 		value := v.Field(i)

// // 		if field.Type.Kind() == reflect.Struct {
// // 			if err := clr.decode(value, unitmap); err != nil {
// // 				return err
// // 			}
// // 			continue
// // 		}

// // 		if field.Type.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Struct {
// // 			if value.IsNil() {
// // 				value.Set(reflect.New(field.Type.Elem()))
// // 			}
// // 			if err := clr.decode(value.Elem(), unitmap); err != nil {
// // 				return err
// // 			}
// // 			continue
// // 		}

// // 		unitName := fmt.Sprintf("%s,%s", field.Tag.Get("full"), field.Tag.Get("short"))
// // 		if unitName == "," {
// // 			unitName = field.Tag.Get("cmd")
// // 			if unitName == "" {
// // 				continue
// // 			} else {
// // 				if _, ok := unitmap[unitName]; ok && !clr.opts.IgnoreCollisions {
// // 					return fmt.Errorf("collision detected for command '%s'", unitName)
// // 				} else {
// // 					unitmap[unitName] = LineUnit{Type: LineUnitTypeCommand, Value: nil}
// // 				}
// // 			}
// // 		}
// // 	}
// // }


// // type CmdConfig struct {
// // 	ActiveCmd  string `cmdpath:true`
// //   Real string `realstring:true`
// //   ExecPath string `execpath:true`
// // 	ConfigPath string `must:true full:"config_path" short:"c" def:"./config.yaml" desc:"Path to config file"`
// // 	RunCmd     struct {
// // 		Port int    `full:"port" short:"p" def:"8080" desc:"Port to run the server on"`
// // 		Host string `full:"host" short:"H" def:"localhost" desc:"Host to run the server on"`
// // 	} `cmd:"run" desc:"Run the server"`
// // }
