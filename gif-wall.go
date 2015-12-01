package main

import (
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/CloudCom/firego"
)

var twitterAPIKey, twitterAPISecret, twitterAccessToken, twitterAccessTokenSecret, twitterStreamFollow string

func loadEnvironmentVariables() {
	twitterAPIKey = os.Getenv("TWITTER_API_KEY")
	if twitterAPIKey == "" {
		log.Fatal("TWITTER_API_KEY not found")
	}
	twitterAPISecret = os.Getenv("TWITTER_API_SECRET")
	if twitterAPIKey == "" {
		log.Fatal("TWITTER_API_SECRET not found")
	}
	twitterAccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	if twitterAPIKey == "" {
		log.Fatal("TWITTER_ACCESS_TOKEN not found")
	}
	twitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	if twitterAPIKey == "" {
		log.Fatal("TWITTER_ACCESS_TOKEN_SECRET not found")
	}
	twitterStreamFollow = os.Getenv("TWITTER_STREAM_FOLLOW")
	if twitterStreamFollow == "" {
		log.Fatal("TWITTER_STREAM_FOLLOW not found")
	}
}

func main() {
	loadEnvironmentVariables()

	//set up firebase connection
	f := firego.New("https://gif-wall.firebaseio.com/gif_list")
	log.Println("Connected to Firebase ...")
	//set up twitter api
	anaconda.SetConsumerKey(twitterAPIKey)
	anaconda.SetConsumerSecret(twitterAPISecret)
	api := anaconda.NewTwitterApi(twitterAccessToken, twitterAccessTokenSecret)
	log.Println("Connected to Twitter Stream API ....")

	v := url.Values{}
	v.Set("follow", twitterStreamFollow)
	stream := api.PublicStreamFilter(v)
	log.Println("Stream initialized ....")
	for t := range stream.C {
		tweet := t.(anaconda.Tweet)
		log.Printf("New tweet : https://twitter.com/%s/status/%s", tweet.User.ScreenName, tweet.IdStr)
		if url := extractGIFFromMediaTwitter(tweet); url != "" {
			saveToFirese(f, url)
		}
		/*if url := extractGIFFromLinkTwitter(tweet); url != "" {
			saveToFirese(f, url)
			continue
		}*/
	}

}
