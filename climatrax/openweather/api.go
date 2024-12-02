package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"climatrax/config"
)

type (
	Weather struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	}

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	}

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	}

	Rain struct {
		OneHour float64 `json:"1h"`
	}

	Clouds struct {
		All int `json:"all"`
	}

	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	}

	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}

	WeatherData struct {
		Coord      Coord     `json:"coord"`
		Weather    []Weather `json:"weather"`
		Base       string    `json:"base"`
		Main       Main      `json:"main"`
		Visibility int       `json:"visibility"`
		Wind       Wind      `json:"wind"`
		Rain       Rain      `json:"rain"`
		Clouds     Clouds    `json:"clouds"`
		DT         int64     `json:"dt"`
		Sys        Sys       `json:"sys"`
		Timezone   int       `json:"timezone"`
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		Cod        int       `json:"cod"`
	}
)

func Query(lat, lon float64) (*WeatherData, error) {
	key := config.GetConfig().GetApiKey()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, key)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	data := &WeatherData{}
	if err := json.Unmarshal(body, data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return data, nil
}
