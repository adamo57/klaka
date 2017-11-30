package db

import (
	"gopkg.in/mgo.v2"
)

//InitDB starts a database connection
func InitDB(port string) (*mgo.Session, error) {
	session, err := mgo.Dial(port)
	if err != nil {
		return nil, err
	}

	if err = session.Ping(); err != nil {
		return nil, err
	}

	return session, nil
}
