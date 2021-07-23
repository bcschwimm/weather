package main

import (
	"fmt"
	"os"
)

type CurrentWeather struct {
	Condition struct {
		Description string `json:"description"`
	} `json:"weather"`

	Detail struct {
		Temp     float64 `json:"temp"`
		Feels    float64 `json:"feels_like"`
		Min      float64 `json:"temp_min"`
		Max      float64 `json:"temp_max"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`

	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunrise"`
	}

	Name string `json:"name"`
}

func main() {
	fmt.Println("ENV:", os.Getenv("WEATHER_KEY"))
}
