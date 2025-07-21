package main

import (
	"fmt"
	"os"
)

func WeatherApiUrl(city string) string {
	if city == "" {
		city = "Saint-Petersburg"
	}
	apiUrl := os.Getenv("WEATHER_API_URL")
	return fmt.Sprintf(apiUrl, city)
}
