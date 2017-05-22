package main

import (
	"regexp"
	"strings"
)

func processTweet(tweet string, config Config) string {
	noRT := regexp.MustCompile("rt @[A-Za-z0-9_]+:? ")
	noNicknames := regexp.MustCompile("@[A-Za-z0-9_]+")
	noLinks := regexp.MustCompile("https?://[^ ]+")
	output := strings.ToLower(tweet)
	output = noRT.ReplaceAllLiteralString(output, "")
	output = noNicknames.ReplaceAllLiteralString(output, "")
	output = noLinks.ReplaceAllLiteralString(output, "")
	for phrase, replacement := range config.Replacements {
		output = strings.Replace(output, phrase, replacement, -1)
	}
	return output
}

func blacklisted(tweet string, config Config) bool {
	for _, blacklist := range config.Blacklist {
		if strings.Contains(tweet, blacklist) {
			return true
		}
	}
	return false
}


func hasReplacement(tweet string, config Config) bool {
	for _, replacement := range config.Replacements {
		if strings.Contains(tweet, replacement) {
			return true
		}
	}
	return false
}

func processTweets(tweets []string, config Config) []string {
	tweetsNew := make(map[string]bool)
	for _, tweet := range tweets {
		output := processTweet(tweet, config)
		if hasReplacement(output, config) && ! blacklisted(output, config) {
			tweetsNew[output] = true
		}
	}
	return map2slice(tweetsNew)
}

func map2slice(m map[string]bool) []string {
	var slice []string
	for k := range m {
		slice = append(slice, k)
	}
	return slice
}
