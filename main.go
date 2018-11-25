package main

import (
	"fmt"
	"log"

	"./keys"
)

func main() {

	api := keys.GetTwitterApi()

	text := "Hello world"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(tweet.Text)
}
