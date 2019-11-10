package main

import (
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/mmcdole/gofeed"
)

func main() {
	feed, _ := gofeed.NewParser().ParseURL("http://b.hatena.ne.jp/hotentry/it.rss")
	api := getTwitterApi()

	for _, item := range feed.Items {
		if !existsTweet(item.Title, api) {
			api.PostTweet(item.Title+"\n"+item.Link, nil)
			return
		}
	}
}

func getTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
}

func existsTweet(title string, api *anaconda.TwitterApi) bool {
	v := url.Values{}
	v.Set("user_id", "1093431019930247168")
	v.Set("count", "10")
	tweets, _ := api.GetUserTimeline(v)
	for _, tweet := range tweets {
		if strings.Contains(tweet.FullText, title) {
			return true
		} else if strings.Contains(strings.Split(tweet.FullText, " - ")[0], strings.Split(title, " - ")[0]) {
			// タイトルURLが含む場合、(title: abcdefg - url)
			// Twitterが自動で短縮するため、「- 」で分割した一件目で比較
			return true
		}
	}
	return false
}
