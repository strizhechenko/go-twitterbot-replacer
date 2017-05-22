package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config is credentials
type Config struct {
	Consumer struct {
		Key    string
		Secret string
	}
	Access struct {
		Token  string
		Secret string
	}
	Replacements map[string]string
	Blacklist map[string]string
}

func parseConfig(path string) Config {
	var config Config
	filename, _ := filepath.Abs(path)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func main() {
	// printTweets(makeTweets())
	webMain()
}
