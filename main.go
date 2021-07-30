package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// apiUrl is our custom string type for the api call
type apiUrl string

// CurrentWeather struct matches the response of our openweather api
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

// DailyWeather matches the onecall api from openweather api
type DailyWeather struct {
	Current struct {
		Sunrise  int     `json:"sunrise"`
		Sunset   int     `json:"sunset"`
		Temp     float64 `json:"temp"`
		Feels    float64 `json:"feels_like"`
		Humidity float64 `json:"humidity"`

		Daily []struct {
			Temp struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
			} `json:"temp"`
		} `json:"daily"`
	} `json:"current"`
}

// populateStruct method on our apiUrl returns
// CurrentWeather with data from the searched city
func (a apiUrl) PopulateStruct() CurrentWeather {

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
func (c CurrentWeather) PrintOutput() {
	fmt.Printf("%s | Currently %.f degrees\n%s\n", c.Name, c.Detail.Temp, strings.Title(c.Condition[0].Description))
	fmt.Printf("Feels like %.f | Minimum Tempature %.f | Maximum Tempature %.f\n", c.Detail.Feels, c.Detail.Min, c.Detail.Max)
	fmt.Printf("Humitidy is at %.f%%\n", c.Detail.Humidity)
	fmt.Printf("Sunrise %v | Sunset %v\n", time.Unix(int64(c.Sys.Sunrise), 0), time.Unix(int64(c.Sys.Sunset), 0))
}

// apiUrlString validates a user input and returns either the api
// call url for a zip code search, or a city name search
func apiUrlString(userInput string) apiUrl {
	_, err := strconv.Atoi(userInput)

	if err != nil {
		return apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=imperial", userInput, os.Getenv("WEATHER_KEY")))
	}
	return apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&appid=%s&units=imperial", userInput, os.Getenv("WEATHER_KEY")))
}

func main() {

	var citySearch = flag.String("c", "philadelphia, pa", "enter a city, state abbreviation or us zipcode to search the weather")
	flag.Parse()
	callUrl := apiUrlString(*citySearch)
	weatherData := callUrl.PopulateStruct()
	weatherData.PrintOutput()
}
