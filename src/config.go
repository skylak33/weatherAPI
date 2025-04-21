package main

import "fmt"

func WeatherApiUrl(city string) string {
	if city == "" {
		city = "Saint-Petersburg"
	}
	return fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&include=current&key=UDZMSJHXEMJK3H293M66NPNCQ&contentType=json", city)
}
