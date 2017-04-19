package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func tweets(c *gin.Context) {
	c.JSON(200, makeTweets())
}

func tweet(c *gin.Context) {
	tweetText := c.PostForm("tweet_text")
	c.JSON(200, gin.H{
		"tweet":  tweetText,
		"result": "OK",
	})
}

func webMain() {
	r := gin.Default()
	r.StaticFile("/", "./assets/index.html")
	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")
	r.GET("/tweets", tweets)
	r.POST("/tweet", tweet)
	r.Run()
}
