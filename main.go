package main

//import "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"

import (
	"os"
	"fmt"
	"log"

	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/mgo.v2"
)

//Handle is what we will want to call when we invoke the lambda job
//func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
//	return "Hello, World!", nil
//}

func main() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		log.Println("error connecting to mongo, make sure you run the mongo instance")
	}

	session.Ping()

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	result, err := api.PostTweet("", nil)

	if err != nil {
		fmt.Println(err)
	}

	log.Println(result.FullText)
}
