package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mom0tomo/jojo-bot/keys"
)

func init() {
	// Twitter 側の問題で golang では http/2 を無効にしないとエラーメッセージが文字化けする。
	godebug := os.Getenv("GODEBUG")
	if godebug != "" {
		godebug += ","
	}
	godebug += "http2client=0"
	os.Setenv("GODEBUG", godebug)
}

func main() {
	api, err := keys.GetTwitterAPI()
	if err != nil {
		log.Fatal(err)
	}

	text := "Hello world"
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(tweet.Text)
}
