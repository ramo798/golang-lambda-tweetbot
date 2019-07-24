package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))

	text := "自作bot test2"

	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(tweet.Text)

}
