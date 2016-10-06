package main

import (
	"log"

	"github.com/CloudCom/firego"
)

func PushToFirebase(f *firego.Firebase, gif interface{}) {
	_, err := f.Push(gif)
	if err != nil {
		log.Printf("Unable to push %s to %s \n", gif, f.String())
	}
	log.Println("Save to Firebase successful !")
}

func UpdateToFirebase(f *firego.Firebase, gif interface{}) {
	err := f.Set(gif)
	if err != nil {
		log.Printf("Unable to update value %s to %s \n", gif, f.String())
	}

	log.Println("Save to Firebase successful !")
}
