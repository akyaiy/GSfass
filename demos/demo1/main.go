package main

import (
	"fmt"
	"sync/atomic"

	"github.com/akyaiy/GSfass/core/config"
)

type AppConfig struct {
	FuncGateway struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"func_gateway"`
	Dashboard struct {
		Enabled bool   `mapstructure:"enabled"`
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
	} `mapstructure:"dashboard"`
}

type App struct {
	Config atomic.Value // holds AppConfig
}

func NewApp() *App {
	app := &App{}
	app.Config.Store(&AppConfig{})
	return app
}

func (a *App) LoadConfig(path string) error {
	err := config.Read().Config().FilePath(path).SetBy(a.Config.Load()).End()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := NewApp()
	err := app.LoadConfig("config/cfg.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded config: %+v\n", app.Config.Load())
}
