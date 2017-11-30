package main

import (
	"log"

	"github.com/adamo57/klaka/db"
	"github.com/adamo57/klaka/twitter"
)

//Handle is what we will want to call when we invoke the lambda job
//func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
//	return "Hello, World!", nil
//}

func main() {
	session, err := db.InitDB(":27017")
	if err != nil {
		log.Fatal("error with the db")
	}

	twitter := twitter.NewTwitterClient()
}
