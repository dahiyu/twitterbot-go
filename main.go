package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
)

func main() {
	// loadEnv()

	// scheduler := gocron.NewScheduler()
	// scheduler.Every(2).Minutes().Do(task)

	api := getTwitterApi()
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("http://b.hatena.ne.jp/hotentry/it.rss")

	for _, item := range feed.Items {
		if !tweeted(item.Title, api) {
			tweet, err := api.PostTweet(item.Title+"\n"+item.Link, nil)
			fmt.Println(err)
			fmt.Println("not tweeted")
			// api.PostTweet(item.Title+"\n"+item.Link, nil)
			fmt.Println(tweet)
			fmt.Println(item.Title + "\n" + item.Link)
			return
		} else {
			fmt.Println("tweeted")
		}
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// 取得したタイトルをすでにツイートしているか判定
func tweeted(title string, api *anaconda.TwitterApi) bool {
	v := url.Values{}
	v.Set("user_id", "1093431019930247168")
	v.Set("count", "100")
	tweets, _ := api.GetUserTimeline(v)
	for _, tweet := range tweets {
		if strings.Contains(tweet.FullText, title) {
			return true
		}
	}
	return false
}
