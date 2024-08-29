package handlers

import (
	"encoding/json"
	"net/http"
	"ums/internal/models"

	"github.com/julienschmidt/httprouter"
)

type (
	UserHandler struct {
		svc *models.Svc
	}
)

// NewUserHandler creates a new UserHandler with the given service
func NewUserHandler(us *models.Svc) *UserHandler {
	return &UserHandler{svc: us}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := uh.svc.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	user, err := uh.svc.GetUserByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := uh.svc.UpdateUser(id, &user); err != nil {
		if err.Error() == "invalid ID format" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if err := uh.svc.DeleteUser(id); err != nil {
		if err.Error() == "invalid ID format" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "User not found", http.StatusNotFound)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
