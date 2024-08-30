package config

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

func Connect() (*mgo.Session, *mgo.Database, error) {
	url := os.Getenv("DB_URL")
	name := os.Getenv("DB_NAME")

	if url == "" || name == "" {
		log.Fatal("DB_URL or DB_NAME environment variable is not set")
	}

	session, err := mgo.Dial(url)
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

	if url == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	session.SetMode(mgo.Monotonic, true)

	log.Println("Connected to MongoDB:", url)
	return session
}
