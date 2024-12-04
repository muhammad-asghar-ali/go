package handlers

import (
	"net/http"

	api "climatrax/openweather"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func Weather(w http.ResponseWriter, r *http.Request) {
	lat_str, lon_str := extractLatLong(r)

	lat, lon, err := validateLatLong(lat_str, lon_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := api.Query(lat, lon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, data)
}
