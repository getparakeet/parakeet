package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	ProjectTitle string
	Key          string
	Language     string
}
func getTomlConfig() tomlConfig {
	var config tomlConfig
	if _, err := toml.DecodeFile("demo.toml", &config); err != nil {
		fmt.Println(err)
	}
	return config
}
func main() {
	fmt.Printf("Initializing project of type %s!", getTomlConfig().Language)
}
