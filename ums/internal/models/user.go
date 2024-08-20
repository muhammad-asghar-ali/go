package models

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID     bson.ObjectId `bson:"_id" json:"_id"`
		Name   string        `bson:"name" json:"name"`
		Gender string        `bson:"gender" json:"gender"`
		Age    int           `bson:"age" json:"age"`
	}

	Svc struct {
		session *mgo.Session
	}
)

// NewSvc creates a new service with the provided session
func NewSvc(session *mgo.Session) *Svc {
	return &Svc{session: session}
}

func (us *Svc) CreateUser(user *User) error {
	user.ID = bson.NewObjectId()

	session := us.session.Clone()
	defer session.Close()

	c := session.DB("ums").C("users")
	return c.Insert(user)
}

func (us *Svc) GetUserByID(id string) (*User, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid ID format")
	}

	session := us.session.Clone()
	defer session.Close()

	c := session.DB("ums").C("users")
	user := &User{}
	if err := c.FindId(bson.ObjectIdHex(id)).One(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *Svc) UpdateUser(id string, updatedUser *User) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid ID format")
	}

	session := us.session.Clone()
	defer session.Close()

	c := session.DB("ums").C("users")
	return c.UpdateId(bson.ObjectIdHex(id), updatedUser)
}

func (us *Svc) DeleteUser(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid ID format")
	}

	session := us.session.Clone()
	defer session.Close()

	c := session.DB("ums").C("users")
	return c.RemoveId(bson.ObjectIdHex(id))
}
