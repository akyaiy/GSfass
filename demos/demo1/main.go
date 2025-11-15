package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/akyaiy/GSfass/core/config"
)

type LiveServer struct {
	current atomic.Value
}

type HTTPserver struct {
	addr   string
	server *http.Server
	ln     net.Listener
}

type FuncGateway struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type AppConfig struct {
	FuncGateway FuncGateway `mapstructure:"func_gateway"`
	Body        string      `mapstructure:"body"`
}

type App struct {
	Config atomic.Value // holds AppConfig
	ls     *LiveServer
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

func (a *App) Server() *LiveServer {
	if a.ls == nil {
		a.ls = &LiveServer{}
	}
	return a.ls
}

func (ls *LiveServer) Start(addr string, handler http.Handler) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := &http.Server{
		Handler: handler,
	}
	hs := &HTTPserver{
		addr:   addr,
		server: srv,
		ln:     ln,
	}
	started := make(chan error, 1)
	go func() {
		err := srv.Serve(ln)
		started <- err
	}()

	select {
	case err := <-started:
		// мгновенная ошибка
		return fmt.Errorf("cannot start server: %w", err)
	case <-time.After(1 * time.Millisecond):
		// если мгновенной ошибки нет — считаем сервер рабочим
	}

	old := ls.current.Load()
	ls.current.Store(hs)
	fmt.Printf("### Old/New server: %+v / %+v\n", old, ls.current.Load())
	if old != nil {
		go func(old *HTTPserver) {
			fmt.Print("### Stopping old server...\n")
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			old.server.Shutdown(ctx)
		}(old.(*HTTPserver))
	}
	return nil
}

func readConfig(conf *AppConfig) {
	for true {
		fmt.Printf("\n\nCurrent config: %+v\n\n", conf)
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	app := NewApp()
	err := app.LoadConfig("config/cfg.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded config: %+v\n", app.Config.Load())

	// ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
	// 	app.Config.Load().(*AppConfig).FuncGateway.Host,
	// 	app.Config.Load().(*AppConfig).FuncGateway.Port,
	// ))
	// if err != nil {
	// 	panic(err)
	// }
	// srv := &http.Server{
	// 	Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		conf := *app.Config.Load().(*AppConfig)
	// 		fmt.Fprintf(w, "%s", conf.Body)
	// 	}),
	// }
	// go srv.Serve(ln)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	conf := *app.Config.Load().(*AppConfig)
	// 	fmt.Fprintf(w, "%s", conf.Body)
	// },
	// )
	// go http.ListenAndServe(fmt.Sprintf("%s:%d",
	// 	app.Config.Load().(*AppConfig).FuncGateway.Host,
	// 	app.Config.Load().(*AppConfig).FuncGateway.Port,
	// ), nil)

	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conf := *app.Config.Load().(*AppConfig)
		fmt.Fprintf(w, "%s", conf.Body)
	})
	go readConfig(app.Config.Load().(*AppConfig))
	actualFuncGateway := app.Config.Load().(*AppConfig).FuncGateway

	if err := app.Server().Start(fmt.Sprintf("%s:%d",
		actualFuncGateway.Host,
		actualFuncGateway.Port,
	), handler); err != nil {
		panic(err)
	} else {
		fmt.Println("### Server started successfully.")
	}
	for true {
		fmt.Scanln()
		err := app.LoadConfig("config/cfg.yaml")
		if err != nil {
			fmt.Printf("### Error reloading config: %v\n", err)
		} else {
			fmt.Println("### Config reloaded successfully.")
		}
		if actualFuncGateway != app.Config.Load().(*AppConfig).FuncGateway {
			fmt.Println("### FuncGateway config changed.")
			actualFuncGateway = app.Config.Load().(*AppConfig).FuncGateway
			if err := app.Server().Start(fmt.Sprintf("%s:%d",
				actualFuncGateway.Host,
				actualFuncGateway.Port,
			), handler); err != nil {
				fmt.Printf("### Error restarting server: %v\n", err)
			} else {
				fmt.Println("### Server restarted successfully.")
			}
		}
	}
}
