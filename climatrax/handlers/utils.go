package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ExtractLatLon(r *http.Request) (string, string) {
	queryParams := r.URL.Query()
	lat := queryParams.Get("lat")
	long := queryParams.Get("lon")

	return lat, long
}

func ValidateLatLon(lat_str, long_str string) (float64, float64, error) {
	if lat_str == "" || long_str == "" {
		return 0, 0, fmt.Errorf("Latitude and longitude are required")
	}

	lat, err := strconv.ParseFloat(lat_str, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid latitude value")
	}

	long, err := strconv.ParseFloat(long_str, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("Invalid longitude value")
	}

	return lat, long, nil
}

func SendJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
