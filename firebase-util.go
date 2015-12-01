package main

import (
	"log"

	"github.com/CloudCom/firego"
)

func saveToFirese(f *firego.Firebase, url string) {
	pushedFirego, err := f.Push(url)
	if err != nil {
		log.Fatal(err)
	}

	var bar string
	if err := pushedFirego.Value(&bar); err != nil {
		log.Fatal(err)
	}
	log.Println("Save to Firebase successful !")
}
