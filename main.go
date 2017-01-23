package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ChimeraCoder/Anaconda"
	"gopkg.in/yaml.v2"
)

var replacements map[string]string

// Config is credentials
type Config struct {
	Consumer     map[string]string
	Access       map[string]string
	Replacements map[string]string
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

func makeAPI(config Config) anaconda.TwitterApi {
	anaconda.SetConsumerKey(config.Consumer["key"])
	anaconda.SetConsumerSecret(config.Consumer["secret"])
	api := anaconda.NewTwitterApi(config.Access["token"], config.Access["secret"])
	return *api
}

func grabTweets(phrase string, api *anaconda.TwitterApi) []anaconda.Tweet {
	tweets, err := api.GetSearch("\""+phrase+"\"", nil)
	if err != nil {
		panic(err)
	}
	return tweets.Statuses
}

func processTweets(tweets []anaconda.Tweet, config Config) {
	var output string
	for _, tweet := range tweets {
		output = tweet.Text
		for phrase, replacement := range config.Replacements {
			output = strings.Replace(output, phrase, replacement, -1)
		}
		for _, replacement := range config.Replacements {
			if strings.Contains(output, replacement) {
				fmt.Println(output)
				break
			}
		}
	}
}

func main() {
	config := parseConfig("./config.yaml")
	api := makeAPI(config)
	for phrase := range config.Replacements {
		tweets := grabTweets(phrase, &api)
		processTweets(tweets, config)
	}
}
