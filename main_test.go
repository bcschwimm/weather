package main

import (
	"fmt"
	"os"
	"testing"
)

func TestApiUrlString(t *testing.T) {

	t.Run("zipcode parsing", func(t *testing.T) {
		got := apiUrlString("90210")
		want := apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=90210&appid=%s&units=imperial", os.Getenv("WEATHER_KEY")))

		if got != want {
			t.Errorf("Error: got %v and want %v", got, want)
		}
	})
	t.Run("city name parsing", func(t *testing.T) {
		got := apiUrlString("philadelphia")
		want := apiUrl(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=philadelphia&appid=%s&units=imperial", os.Getenv("WEATHER_KEY")))

		if got != want {
			t.Errorf("Error: got %v and want %v", got, want)
		}
	})
}
