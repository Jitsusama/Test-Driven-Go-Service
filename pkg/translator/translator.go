package translator

// Translator is able to translate temperature requests to/from weather requests.
type Translator struct {
	weather WeatherRetriever
}

// Create a new Translator.
func Create(weather WeatherRetriever) *Translator {
	return &Translator{weather}
}

// RetrieveTemperature from a WeatherRetriever.
func (w *Translator) RetrieveTemperature(city string) (string, error) {
	weather, _ := w.weather.RetrieveWeather(city, "j1")
	return weather.CurrentTempInCelsius, nil
}
