package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

func Connect() (*mgo.Session, *mgo.Database, error) {
	url := os.Getenv("DB_URL")
	name := os.Getenv("DB_NAME")

	session, err := mgo.DialWithTimeout(url, 10*time.Second)
	if err != nil {
		return nil, nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	database := session.DB(name)

	log.Println("Connected to MongoDB:", url)

	return session, database, nil
}

func GetSession() *mgo.Session {
	url := os.Getenv("DB_URL")

	session, err := mgo.DialWithTimeout(url, 10*time.Second)
	if err != nil {
		return nil
	}

	session.SetMode(mgo.Monotonic, true)

	log.Println("Connected to MongoDB:", url)
	return session
}
