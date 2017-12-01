package twitter

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

//NewTwitterClient creates a twitter client
func NewTwitterClient(accessToken string, accessTokenSecret string) *anaconda.TwitterApi {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}
