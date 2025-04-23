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
}

type WeatherResponse struct {
	Address  string       `json:"address"`
	Timezone string       `json:"timezone"`
	Days     []WeatherDay `json:"days"`
}
type OutputData struct {
	City     string  `json:"city"`
	Datetime string  `json:"datetime"`
	Temp     float64 `json:"temp"`
}

func main() {
	router := gin.Default()
	router.GET("/weather", getWeather)
	router.Run(":8080")
}

func getWeather(c *gin.Context) {

	city := c.Query("city")

	resp, err := http.Get(WeatherApiUrl(city))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	if len(weather.Days) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No weather data found"})
		return
	}

	answer := OutputData{
		City:     weather.Address,
		Datetime: weather.Days[0].Datetime,
		Temp:     weather.Days[0].Temp,
	}
	c.IndentedJSON(http.StatusOK, answer)
}
