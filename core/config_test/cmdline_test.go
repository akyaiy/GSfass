package config

// import (
// 	"testing"
// 	//"github.com/akyaiy/GSfass/core/config"
// )

// // type CmdConfig struct {
// // 	Foo string `full:"foo" short:"f" def:"defaultFoo" desc:"Foo flag"`
// // 	Bar int    `full:"bar" short:"b" def:"42" desc:"Bar flag"`
// // }

// type CmdConfig struct {
// 	ActiveCmd  string `cmdpath:true`
// 	ConfigPath string `must:true full:"config_path" short:"c" def:"./config.yaml" desc:"Path to config file"`
// 	RunCmd     struct {
// 		Port int    `full:"port" short:"p" def:"8080" desc:"Port to run the server on"`
// 		Host string `full:"host" short:"H" def:"localhost" desc:"Host to run the server on"`
// 	} `cmd:"run" desc:"Run the server"`
// }

// func Test_cmdLineReader(t *testing.T) {

// }
