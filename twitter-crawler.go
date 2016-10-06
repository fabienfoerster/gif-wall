package main

import (
	"log"
	"mime"
	"path/filepath"

	"github.com/ChimeraCoder/anaconda"
)

type GIF struct {
	URL      string `json:"url"`
	Mimetype string `json:"mimetype"`
}

func extractGIFFromMediaTwitter(tweet anaconda.Tweet) GIF {

	if medias := tweet.ExtendedEntities.Media; len(medias) != 0 {
		if variants := medias[0].VideoInfo.Variants; len(variants) != 0 {
			url := variants[0].Url
			mimetype := mime.TypeByExtension(filepath.Ext(url))
			log.Printf("Media GIF found : %s", url)
			gif := GIF{url, mimetype}
			return gif
		}
	}
	log.Println("No media GIF found")
	return GIF{}

}

func extractGIFFromLinkTwitter(tweet anaconda.Tweet) GIF {
	if urls := tweet.Entities.Urls; len(urls) != 0 {
		for _, url := range urls {
			expendedURL := url.Expanded_url
			if extension := filepath.Ext(expendedURL); extension == ".gif" {
				log.Printf("GIF link found : %s", expendedURL)
				mimetype := mime.TypeByExtension(filepath.Ext(expendedURL))
				gif := GIF{expendedURL, mimetype}
				return gif
			}
		}

	}
	log.Println("No GIF link found")
	return GIF{}
}
