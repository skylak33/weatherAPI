package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	//fmt.Println("WEATHER_API_URL:", os.Getenv("WEATHER_API_URL"))
	router := gin.Default()
	router.GET("/weather", getWeather)
	router.Run(":8080")
}
