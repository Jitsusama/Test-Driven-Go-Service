package translator

// WeatherRetriever retrieves weather from the wttr service.
type WeatherRetriever interface {
	// RetrieveWeather for a city in a certain format.
	RetrieveWeather(city string, format string) (Weather, error)
}

// Weather represents the weather for a certain location.
type Weather struct {
	CurrentTempInCelsius string
}
