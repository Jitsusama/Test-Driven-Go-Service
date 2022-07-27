package translator

// TemperatureRetriever retrieves the temperature.
type TemperatureRetriever interface {
	RetrieveTemperature(city string) (string, error)
}
