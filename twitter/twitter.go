package twitter

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

//NewTwitterClient creates a twitter client
func NewTwitterClient() *anaconda.TwitterApi {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}
