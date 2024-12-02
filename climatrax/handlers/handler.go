package handlers

import (
	"net/http"

	api "climatrax/openweather"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func Weather(w http.ResponseWriter, r *http.Request) {
	latStr, lonStr := ExtractLatLon(r)

	lat, long, err := ValidateLatLon(latStr, lonStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := api.Query(lat, long)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendJSONResponse(w, data)
}
