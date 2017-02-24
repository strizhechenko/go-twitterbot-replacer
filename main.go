package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ChimeraCoder/Anaconda"
	"gopkg.in/yaml.v2"
)

var replacements map[string]string

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

func makeAPI(config Config) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(config.Consumer.Key)
	anaconda.SetConsumerSecret(config.Consumer.Secret)
	api := anaconda.NewTwitterApi(config.Access.Token, config.Access.Secret)
	return api
}

func grabTweets(phrase string, api *anaconda.TwitterApi) []anaconda.Tweet {
	tweets, err := api.GetSearch("\""+phrase+"\"", nil)
	if err != nil {
		panic(err)
	}
	return tweets.Statuses
}

func processTweets(tweets []anaconda.Tweet, config Config) map[string]bool {
	var output string
	noRT := regexp.MustCompile("rt @[A-Za-z0-9_]+:? ")
	noNicknames := regexp.MustCompile("@[A-Za-z0-9_]+")
	noLinks := regexp.MustCompile("https?://[^ ]+")
	tweetsNew := make(map[string]bool)
	for _, tweet := range tweets {
		output = tweet.Text
		for phrase, replacement := range config.Replacements {
			output = strings.Replace(output, phrase, replacement, -1)
			output = strings.ToLower(output)
			output = noRT.ReplaceAllLiteralString(output, "")
			output = noNicknames.ReplaceAllLiteralString(output, "")
			output = noLinks.ReplaceAllLiteralString(output, "")
		}
		for _, replacement := range config.Replacements {
			if strings.Contains(output, replacement) {
				tweetsNew[output] = true
				break
			}
		}
	}
	return tweetsNew
}

func printTweets(tweets map[string]bool) {
	for tweet := range tweets {
		fmt.Println(tweet)
	}
}

func main() {
	config := parseConfig("./config.yaml")
	api := makeAPI(config)
	for phrase := range config.Replacements {
		tweets := grabTweets(phrase, api)
		tweetsNew := processTweets(tweets, config)
		printTweets(tweetsNew)
	}
}
