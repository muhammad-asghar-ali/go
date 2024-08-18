package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

type (
	UserHandler struct {
		s *mgo.Session
	}
)

func NewUserHandler(s *mgo.Session) *UserHandler {
	return &UserHandler{s}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
