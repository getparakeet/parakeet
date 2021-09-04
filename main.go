package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
func checkTomlConfig() {
	key := getTomlConfig().Key
	title := getTomlConfig().ProjectTitle
	if key == "" {
		panic("No key found in parakeet.toml")
	} else if strings.Contains(key, "<script>") {
		panic("You're not allowed to do that. Stop being naughty.")
	} else if strings.Contains(key, "SELECT") {
		panic("You're not allowed to do that. Stop being naughty.")
	} else if strings.Contains(key, "DROP") {
		panic("You're not allowed to do that. Stop being naughty.")
	}
	body, err := json.Marshal(map[string]string{
		"projectTitle": title,
		"key":          key,
	})
	if err != nil {
		panic(err)
	}
	res, err := http.Post("https://api.parakeet.cloud/v1/verify/key", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 400 {
		panic("Invalid key. Please check `Key` and `ProjectTitle` in your parakeet.toml file.")
	}
}
func main() {
	fmt.Printf("Initializing project of type %s\n", getTomlConfig().Language)
	checkTomlConfig()
	
}
