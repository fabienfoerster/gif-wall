package main

import (
	"log"
	"path/filepath"

	"github.com/ChimeraCoder/anaconda"
)

func extractGIFFromMediaTwitter(tweet anaconda.Tweet) string {

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

func extractGIFFromLinkTwitter(tweet anaconda.Tweet) string {
	if urls := tweet.Entities.Urls; len(urls) != 0 {
		for _, url := range urls {
			expendedURL := url.Expanded_url
			if extension := filepath.Ext(expendedURL); extension == ".gif" {
				log.Printf("GIF link found : %s", expendedURL)
				return expendedURL
			}
		}

	}
	log.Println("No GIF link found")
	return ""
}
