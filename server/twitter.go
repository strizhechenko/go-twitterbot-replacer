package main

import (
	"fmt"

	"github.com/ChimeraCoder/Anaconda"
)

func makeAPI(config Config) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(config.Consumer.Key)
	anaconda.SetConsumerSecret(config.Consumer.Secret)
	api := anaconda.NewTwitterApi(config.Access.Token, config.Access.Secret)
	return api
}

func grabTweets(phrase string, api *anaconda.TwitterApi) []string {
	var statuses []string
	tweets, err := api.GetSearch("\""+phrase+"\"", nil)
	if err != nil {
		panic(err)
	}
	for _, v := range tweets.Statuses {
		statuses = append(statuses, v.Text)
	}
	return statuses
}

func makeTweets() []string {
	config := parseConfig("./config.yaml")
	newTweets := make(map[string]bool)
	api := makeAPI(config)
	for phrase := range config.Replacements {
		tweets := grabTweets(phrase, api)
		tweetsNew := processTweets(tweets, config)
		for _, v := range tweetsNew {
			newTweets[v] = true
		}
	}
	api.Close()
	return map2slice(newTweets)
}

func printTweets(tweets []string) {
	for _, tweet := range tweets {
		fmt.Println(tweet)
	}
}
