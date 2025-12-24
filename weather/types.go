package weather

type WeatherResponseDescription struct {
	Id          int32  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherResponseDaily struct {
	DT        int32   `json:"dt"`
	Sunrise   int32   `json:"sunrise"`
	Sunset    int32   `json:"sunset"`
	Moonrise  int32   `json:"moonrise"`
	Moonset   int32   `json:"moonset"`
	MoonPhase float32 `json:"moon_phase"`
	Summary   string  `json:"summary"`
	Temp      struct {
		Day   float32 `json:"day"`
		Min   float32 `json:"min"`
		Max   float32 `json:"max"`
		Night float32 `json:"night"`
		Eve   float32 `json:"eve"`
		Morn  float32 `json:"morn"`
	} `json:"temp"`
	FeelsLike struct {
		Day   float32 `json:"day"`
		Night float32 `json:"night"`
		Eve   float32 `json:"eve"`
		Morn  float32 `json:"morn"`
	} `json:"feels_like"`
	Pressure  int32                        `json:"pressure"`
	Humidity  int32                        `json:"humidity"`
	DewPoint  float32                      `json:"dew_point"`
	WindSpeed float32                      `json:"wind_speed"`
	WindDeg   float32                      `json:"wind_deg"`
	WindGust  float32                      `json:"wind_gust"`
	Weather   []WeatherResponseDescription `json:"weather"`
	Clouds    int32                        `json:"clouds"`
	POP       float32                      `json:"pop"`
	Rain      float32                      `json:"rain"`
	UVI       float32                      `json:"uvi"`
}

type WeatherResponseHourly struct {
	DT         int32                        `json:"dt"`
	Temp       float32                      `json:"temp"`
	FeelsLike  float32                      `json:"feels_like"`
	Pressure   int32                        `json:"pressure"`
	Humidity   int32                        `json:"humidity"`
	DewPoint   float32                      `json:"dew_point"`
	UVI        float32                      `json:"uvi"`
	Clouds     int32                        `json:"clouds"`
	Visibility int32                        `json:"visibility"`
	WindSpeed  float32                      `json:"wind_speed"`
	WindDeg    float32                      `json:"wind_deg"`
	WindGust   float32                      `json:"wind_gust"`
	Weather    []WeatherResponseDescription `json:"weather"`
	POP        int32                        `json:"pop"`
}

type WeatherResponseCurrent struct {
	DT         int32                        `json:"dt"`
	Sunrise    int32                        `json:"sunrise"`
	Sunset     int32                        `json:"sunset"`
	Temp       float32                      `json:"temp"`
	FeelsLike  float32                      `json:"feels_like"`
	Pressure   int32                        `json:"pressure"`
	Humidity   int32                        `json:"humidity"`
	DewPoint   float32                      `json:"dew_point"`
	UVI        float32                      `json:"uvi"`
	Clouds     int32                        `json:"clouds"`
	Visibility int32                        `json:"visibility"`
	WindSpeed  float32                      `json:"wind_speed"`
	WindDeg    float32                      `json:"wind_deg"`
	Weather    []WeatherResponseDescription `json:"weather"`
}

type WeatherResponse struct {
	Lat     float64                 `json:"lat"`
	Lon     float64                 `json:"lon"`
	Current WeatherResponseCurrent  `json:"current"`
	Hourly  []WeatherResponseHourly `json:"hourly"`
	Daily   []WeatherResponseDaily  `json:"daily"`
}
