package main

import (
	"log"

	"github.com/adamo57/klaka/db"
	"gopkg.in/mgo.v2/bson"
)

//Handle is what we will want to call when we invoke the lambda job
//func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
//	return "Hello, World!", nil
//}

//TwitterRecord holds the data for the twitter entries
type TwitterRecord struct {
	Twitter TwitterUser `bson:"twitter"`
}

//TwitterUser holds the data for the twitter entries
type TwitterUser struct {
	OAuthToken       string `bson:"oauth_token"`
	OAuthTokenSecret string `bson:"oauth_token_secret"`
	UserID           string `bson:"user_id"`
	ScreenName       string `bson:"screen_name"`
}

func main() {
	//for testing purposes only
	findList := []string{"adamo57"}

	s, err := db.InitDB(":27017")
	if err != nil {
		log.Fatal("error with the db")
	}

	c := s.DB("test").C("twitter")

	var results TwitterRecord

	//> db.twitter.find({"twitter.screen_name" : "adamo57"})
	for _, value := range findList {
		err = c.Find(bson.M{"twitter.screen_name": value}).One(&results)
		if err != nil {
			log.Fatal(err)
		}
	}
}
