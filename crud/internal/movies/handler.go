package movies

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	dummy()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	dummy()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	dummy()
	w.Header().Set("Content-Type", "application/json")
	movie := &Movie{}

	if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	dummy()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)

			movie := &Movie{}
			if err := json.NewDecoder(r.Body).Decode(movie); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}
			movie.ID = strconv.Itoa(rand.Intn(1000000000))
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	dummy()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}
