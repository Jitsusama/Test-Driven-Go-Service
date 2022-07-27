package translator

type Translator interface {
	RetrieveWeather(city string) string
}
