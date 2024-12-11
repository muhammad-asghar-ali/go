package models

import (
	"pulse/internal/errors"
)

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (u *User) FindUserByID(id int, users []User) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, errors.ErrUserNotFoundInProducer
}
