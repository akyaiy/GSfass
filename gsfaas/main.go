package main

import (
	"github.com/akyaiy/GSfass/core/config"
)

func main() {
	var env struct {
		Port string `mapstructure:"PORT"`
		Host string `mapstructure:"HOST"`
	}
	if err := config.Read().Environment().SetBy(&env).SetDefaults(map[string]string{
		"PORT": "8080",
	}).End(); err != nil {
		panic(err)
	}

	println("Host:", env.Host)
	println("Port:", env.Port)
}
