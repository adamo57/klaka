package main

//import "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	"fmt"
	"net/url"
	"log"
)

func init() {

}

//Handle is what we will want to call when we invoke the lambda job
//func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
//	return "Hello, World!", nil
//}

func main() {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	result, err := api.PostTweet("This is posted with <3 from golang", nil)

	if err != nil {
		fmt.Println(err)
	}

	log.Println(result.FullText)
}
