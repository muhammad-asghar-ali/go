package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func extractLatLong(r *http.Request) (string, string) {
	queryParams := r.URL.Query()
	lat := queryParams.Get("lat")
	lon := queryParams.Get("lon")

	return lat, lon
}

func validateLatLong(lat_str, lon_str string) (float64, float64, error) {
	if lat_str == "" || lon_str == "" {
		return 0, 0, fmt.Errorf("Latitude and longitude are required")
	}

	lat, err := strconv.ParseFloat(lat_str, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid latitude value")
	}

	lon, err := strconv.ParseFloat(lon_str, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid longitude value")
	}

	return lat, lon, nil
}

func sendJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
