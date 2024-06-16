package types

type WeatherInfo struct {
	Location Location           `json:"location"`
	Current  CurrentWeatherInfo `json:"current"`
}

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TzId      string  `json:"tz_id"`
	LocalTime string  `json:"localtime"`
}

type CurrentWeatherInfo struct {
	TemperatureInF float32 `json:"temp_f"`
	TemperatureInC float32 `json:"temp_c"`

	LastUpdateTime string `json:"last_updated"`

	Condition Condition `json:"condition"`

	WindSpeedKph     float32 `json:"wind_kph"`
	WindGustSpeedMph float32 `json:"wind_gust_mph"`
	WindDegrees      float32 `json:"wind_degree"`
	WindDirection    string  `json:"wind_dir"`

	PressureMb float32 `json:"pressure_mb"`
	PressureIn float32 `json:"pressure_in"`

	PrecipIn float32 `json:"precip_in"`
	PrecipMm float32 `json:"precip_mm"`

	Humidity float32 `json:"humidity"`

	CloudCover float32 `json:"cloud"`

	FeelsLikeC float32 `json:"feelslike_c"`
	FeelslikeF float32 `json:"feelslike_f"`

	WindchillC float32 `json:"windchill_c"`
	WindchillF float32 `json:"windchill_f"`

	HeatIndexC float32 `json:"heatindex_c"`
	HeatIndexF float32 `json:"heatindex_f"`

	DewPointC float32 `json:"dew_point_c"`
	DewPointF float32 `json:"dew_point_f"`

	VisKm    float32 `json:"vis_km"`
	VisMiles float32 `json:"vis_miles"`

	Uv float32 `json:"uv"`

	GustMph float32 `json:"gust_mph"`
	GustKph float32 `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}
