package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	checker()

}

func makesentence(rainflag bool, stormflag bool, snowflag bool) {
	var sentence string

	// テスト用
	// rainflag = true

	sentence += "【BOT】"

	t := time.Now()
	const layout = "本日、15:04:05 JST 時点の和歌山市で"
	// fmt.Println(reflect.TypeOf(t.Format(layout)))
	// fmt.Println(t.Format(layout))
	sentence += t.Format(layout)

	if rainflag {
		sentence += "大雨警報 "
	}

	if stormflag {
		sentence += "暴風警報 "
	}

	if snowflag {
		sentence += "大雪警報 "
	}

	if rainflag || stormflag || rainflag {
		sentence += "が発令されています。"
		tweet(sentence)
	} else {
		fmt.Println("警報なし")
	}

	// fmt.Println(sentence)
}

func checker() {
	doc, err := goquery.NewDocument("https://www.jma.go.jp/jp/warn/336_table.html")
	if err != nil {
		panic(err)
	}

	selection := doc.Find("#WarnTableTable > tbody > tr").First()
	targettr := selection.Next().Next() //tr:nth-child(3)

	innertext := targettr.Find("td").First().Next().Next() //td:nth-child(3)

	var rainflag = false
	var stormflag = false
	var snowflag = false

	//大雨警報の判定　#WarnTableTable > tbody > tr:nth-child(3) > td:nth-child(4)
	rain := innertext.Next()
	if rain.Text() == "●" {
		rainflag = true
	}
	// fmt.Println(rainflag)

	//暴風警報の判定　#WarnTableTable > tbody > tr:nth-child(3) > td:nth-child(6)
	storm := innertext.Next().Next().Next()
	if storm.Text() == "●" {
		stormflag = true
	}
	// fmt.Println(stormflag)

	//大雪警報の判定　#WarnTableTable > tbody > tr:nth-child(3) > td:nth-child(8)
	snow := innertext.Next().Next().Next().Next().Next()
	if snow.Text() == "●" {
		snowflag = true
	}
	// fmt.Println(snowflag)

	//テストコード　td:nth-child(4)
	// var testflag = false
	// test := innertext.Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next().Next()
	// fmt.Println(test.Text())
	// if test.Text() == "●" {
	// 	testflag = true
	// }
	// if testflag {
	// 	fmt.Println("成功")
	// }
	//typeはstring
	// fmt.Println(reflect.TypeOf(test.Text()))

	makesentence(rainflag, stormflag, snowflag)
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
