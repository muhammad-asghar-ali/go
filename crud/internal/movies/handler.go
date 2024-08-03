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

	_ = json.NewDecoder(r.Body).Decode(movie)

	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMoive(w http.ResponseWriter, r *http.Request) {
	dummy()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)

			movie := &Movie{}
			_ = json.NewDecoder(r.Body).Decode(movie)
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
