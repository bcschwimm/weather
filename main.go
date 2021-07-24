package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// CurrentWeather matches the return of our open
// weather api
type CurrentWeather struct {
	Condition []struct {
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
		Sunset  int `json:"sunset"`
	}

	Name string `json:"name"`
}

///// define print method for struct
///// create & define api call method for api string custom type, have method return w.

type apiUrl string

// apiUrlString validates a user input and returns either the api
// call url for a zip code search, or a city name search
func apiUrlString(userInput string) apiUrl {
	_, err := strconv.Atoi(userInput)

	if err != nil {
		return apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", userInput, os.Getenv("WEATHER_KEY")))
	}
	return apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s", userInput, os.Getenv("WEATHER_KEY")))
}

// apiCall takes in a search city name and returns
// a populated CurrentWeather struct with weather data
func apiCall(cityName string) CurrentWeather {

	s := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityName, os.Getenv("WEATHER_KEY"))

	resp, err := http.Get(s)

	if err != nil {
		fmt.Printf("Error with API Request %s\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error Reading Response Body %s\n", err)
	}

	var weatherData CurrentWeather
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		fmt.Printf("Error unmarshalling json body into WeatherData struct %s\n", err)
	}

	return weatherData
}

// implement flag arg for city, with the default being our current location
func main() {

	var citySearch = flag.String("c", "philadelphia, pa", "enter a city, state abbreviation to search the weather")
	flag.Parse()
	w := apiCall(*citySearch)

	fmt.Println(w.Name)
}
