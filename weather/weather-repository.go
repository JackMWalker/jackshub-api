package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"jackshub/api/logger"
	"net/http"
)

const ApiKey = "b52052066ccc597f4e6ad510bb76177f"
const Lat = 52.619271 // tamworth
const Lon = -1.639175 // tamworth
const Units = "metric"
const Exclude = "minutely"
const BaseUri = "https://api.openweathermap.org/data/3.0/onecall?lat=%v&lon=%v&exclude=%s&units=%s&appid=%s"

func GetForecast() (*WeatherResponse, error) {
	logger := logger.Logger{Context: "weather-repository"}

	url := fmt.Sprintf(BaseUri, Lat, Lon, Exclude, Units, ApiKey)

	logger.Log("making request to [%s]", url)

	resp, err := http.Get(url)

	if err != nil {
		logger.Error("error making http request", err)
		return nil, err
	}

	logger.Log("got response from [%s] status code %d", url, resp.StatusCode)

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("could not read response body", err)
		return nil, err
	}

	var weatherResponse WeatherResponse

	err = json.Unmarshal(resBody, &weatherResponse)

	if err != nil {
		logger.Error("unexpected response body structure", err)
		return nil, err
	}

	logger.Log("response body: %s\n", resBody)

	return &weatherResponse, nil
}
