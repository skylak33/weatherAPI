package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherDay struct {
	Datetime      string  `json:"datetime"`
	DatetimeEpoch int64   `json:"datetimeEpoch"`
	Temp          float64 `json:"temp"`
	Feelslike     float64 `json:"feelslike"`
	Windspeed     float64 `json:"windspeed"`
	Condition     string  `json:"conditions"`
}

type WeatherResponse struct {
	Address  string       `json:"address"`
	Timezone string       `json:"timezone"`
	Days     []WeatherDay `json:"days"`
}
type OutputData struct {
	City          string    `json:"city"`
	Datetime      []string  `json:"datetime"`
	DatetimeEpoch []int64   `json:"datetimeEpoch"`
	Temp          []float64 `json:"temp"`
	Feelslike     []float64 `json:"feelslike"`
	Windspeed     []float64 `json:"windspeed"`
	Condition     []string  `json:"conditions"`
}

func extractField[T any](days []WeatherDay, extractor func(WeatherDay) T) []T {
	result := make([]T, len(days))
	for i, day := range days {
		result[i] = extractor(day)
	}
	return result
}

func getWeather(c *gin.Context) {
    city := c.Query("city")
    escapedCity := url.QueryEscape(city)
    url := WeatherApiUrl(escapedCity)
    fmt.Println("Request URL:", url)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("HTTP error:", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Read body error:", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
        return
    }

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("API error (status %d): %s\n", resp.StatusCode, string(body))
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Weather API error"})
        return
    }

    var weather WeatherResponse
    if err := json.Unmarshal(body, &weather); err != nil {
        fmt.Println("JSON parse error:", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
        return
    }
    answer := OutputData{
        City:          weather.Address,
        Datetime:      extractField(weather.Days, func(day WeatherDay) string { return day.Datetime }),
        DatetimeEpoch: extractField(weather.Days, func(day WeatherDay) int64 { return day.DatetimeEpoch }),
        Temp:          extractField(weather.Days, func(day WeatherDay) float64 { return day.Temp }),
        Feelslike:     extractField(weather.Days, func(day WeatherDay) float64 { return day.Feelslike }),
        Windspeed:     extractField(weather.Days, func(day WeatherDay) float64 { return day.Windspeed }),
        Condition:     extractField(weather.Days, func(day WeatherDay) string { return day.Condition }),
    }
    c.IndentedJSON(http.StatusOK, answer)
}
