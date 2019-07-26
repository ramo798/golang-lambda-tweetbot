package main

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	text := "あ"
	tweet(text)
}

func tweet(tweettext string) {
	anaconda.SetConsumerKey(os.Getenv("TWITTERCONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTERCONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTERACCESS_TOKEN"), os.Getenv("TWITTERACCESS_TOKEN_SECRET"))

	tweet, err := api.PostTweet(tweettext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(tweet.Text)

}
