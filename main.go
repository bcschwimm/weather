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

// apiUrl is our custom string type for the api call
type apiUrl string

// CurrentWeather struct matches the response of our open weather api
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

// populateStruct method on our apiUrl returns
// CurrentWeather with data from the searched city
func (a apiUrl) populateStruct() CurrentWeather {

	// converting apiUrl back to string for http get
	resp, err := http.Get(string(a))

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

// printOutput method on our CurrentWeather struct formats
// the response and prints it to the terminal
func (c CurrentWeather) printOutput() {
	fmt.Println(c.Name, c.Detail.Temp)
}

// apiUrlString validates a user input and returns either the api
// call url for a zip code search, or a city name search
func apiUrlString(userInput string) apiUrl {
	_, err := strconv.Atoi(userInput)

	if err != nil {
		return apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", userInput, os.Getenv("WEATHER_KEY")))
	}
	return apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s", userInput, os.Getenv("WEATHER_KEY")))
}

func main() {

	var citySearch = flag.String("c", "philadelphia, pa", "enter a city, state abbreviation or us zipcode to search the weather")
	flag.Parse()

	call := apiUrlString(*citySearch)
	w := call.populateStruct()
	w.printOutput()
}
