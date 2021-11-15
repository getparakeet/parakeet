package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/getparakeet/parakeet/errors"
	"github.com/getparakeet/parakeet/src"
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
		errors.NoKeyError()
	} else if strings.Contains(key, "<script>") {
		errors.SecurityError(fmt.Errorf("key contains <script> tag"))
	} else if strings.Contains(key, "SELECT") {
		errors.SecurityError(fmt.Errorf("key contains SELECT statement"))
	} else if strings.Contains(key, "DROP") {
		errors.SecurityError(fmt.Errorf("key contains DROP statement"))
	}
	body, err := json.Marshal(map[string]string{
		"projectTitle": title,
		"key":          key,
	})
	if err != nil {
		errors.UnknownError(err)
	}
	res := src.PostHttp("https://api.parakeet.cloud/v1/verify/key", body)
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		errors.UnknownError(err)
	}
	if res.StatusCode == 400 {
		errors.UnknownError(fmt.Errorf("invalid key. Please check `Key` and `ProjectTitle` in your parakeet.toml file"))
	}
}
func main() {
	fmt.Printf("Initializing project of type %s\n", getTomlConfig().Language)
	checkTomlConfig()
}
