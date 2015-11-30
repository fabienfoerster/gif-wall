package main

import (
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/CloudCom/firego"
)

func main() {

	//set up firebase connection
	f := firego.New("https://gif-wall.firebaseio.com/gif_list")
	log.Println("Connected to Firebase ...")
	//set up twitter api
	anaconda.SetConsumerKey("2vxn8zOFZuBPCep3oejC33SsX")
	anaconda.SetConsumerSecret("8mjb3gIK7agiZYvEiQQHFjxDbqMMzqCzrbNXJg0aDkTyHt3NIk")
	api := anaconda.NewTwitterApi("1620076356-iQfWV1al7DCa3XVbfl3mLsmAS2GgCEB22WFtGgu", "lbhjpDSZ7Qx0fcG75P9YMC7sLB0UTlnSF8OAminwOyRZU")
	log.Println("Connected to Twitter Stream API ....")
	v := url.Values{}
	v.Set("follow", "168749152")
	stream := api.PublicStreamFilter(v)
	log.Println("Stream initialized ....")
	for t := range stream.C {
		tweet := t.(anaconda.Tweet)
		log.Printf("New tweet : https://twitter.com/%s/status/%s", tweet.User.ScreenName, tweet.IdStr)
		if url := extractGifFromMediaTwitter(tweet); url != "" {
			pushedFirego, err := f.Push(url)
			if err != nil {
				log.Fatal(err)
			}

			var bar string
			if err := pushedFirego.Value(&bar); err != nil {
				log.Fatal(err)
			}
		}
	}

}

func extractGifFromMediaTwitter(tweet anaconda.Tweet) string {

	if medias := tweet.ExtendedEntities.Media; len(medias) != 0 {
		if variants := medias[0].VideoInfo.Variants; len(variants) != 0 {
			url := variants[0].Url
			log.Printf("Media GIF found : %s", url)
			return url
		}
	}
	log.Println("No media GIF found")
	return ""

}
