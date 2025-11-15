package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/akyaiy/GSfass/core/config"
)

type AppConfig struct {
	FuncGateway struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"func_gateway"`
	Body string `mapstructure:"body"`
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

func readConfig(conf *AppConfig) {
	for true {
		fmt.Printf("\n\nCurrent config: %+v\n\n", conf)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	app := NewApp()
	err := app.LoadConfig("config/cfg.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded config: %+v\n", app.Config.Load())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			conf := app.Config.Load().(*AppConfig)
			fmt.Fprintf(w, "%s" ,conf.Body)
		},
	)
	go http.ListenAndServe(fmt.Sprintf("%s:%d",
		app.Config.Load().(*AppConfig).FuncGateway.Host,
		app.Config.Load().(*AppConfig).FuncGateway.Port,
	), nil)

	go readConfig(app.Config.Load().(*AppConfig))
	for true {
		fmt.Scanln()
		err := app.LoadConfig("config/cfg.yaml")
		if err != nil {
			fmt.Printf("### Error reloading config: %v\n", err)
		} else {
			fmt.Println("### Config reloaded successfully.")
		}
	}
}
