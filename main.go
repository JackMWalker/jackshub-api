package main

import (
	"jackshub/api/logger"
	"jackshub/api/weather"
	"os"
)

func main() {
	logger := logger.Logger{Context: "main"}

	weatherData, err := weather.GetForecast()

	if err != nil {
		logger.Log("error retrieving weather data")
		os.Exit(1)
	}

	logger.Log("Current cloud cover: %d", weatherData.Current.Clouds)
}
