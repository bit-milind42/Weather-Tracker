package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

type apiConfigData struct {
	OpenWeatherMapAPIKey string `json:"openWeatherMapAPIKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData
	return c, json.Unmarshal(bytes, &c)
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s",
		apiConfig.OpenWeatherMapAPIKey, url.QueryEscape(city))
	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return weatherData{}, err
	}

	var d weatherData
	return d, json.Unmarshal(body, &d)
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, from Go!\n"))
	})

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 {
			http.Error(w, "City name is required in URL path", http.StatusBadRequest)
			return
		}
		city := pathParts[2]
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
