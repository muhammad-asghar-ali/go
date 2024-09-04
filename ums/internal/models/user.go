package models

import (
	"context"
	"errors"
	"ums/internal/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
		Name   string             `bson:"name" json:"name"`
		Gender string             `bson:"gender" json:"gender"`
		Age    int                `bson:"age" json:"age"`
	}

	Svc struct {
		client *mongo.Client
	}
)

// NewSvc creates a new service with the provided session
func NewSvc(client *mongo.Client) *Svc {
	return &Svc{client: client}
}

func (us *Svc) CreateUser(user *User) error {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	collection := config.GetUserCollection(us.client)
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (us *Svc) GetUserByID(id string) (*User, error) {
	collection := config.GetUserCollection(us.client)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	user := &User{}

	err = collection.FindOne(context.Background(), filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *Svc) UpdateUser(id string, u *User) error {
	collection := config.GetUserCollection(us.client)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": u}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	return err
}

func (us *Svc) DeleteUser(id string) error {
	collection := config.GetUserCollection(us.client)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)

	return err
}
