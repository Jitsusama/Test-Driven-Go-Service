package wttr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"playing-around/pkg/translator"
)

type Wttr struct {
	baseUri string
}

func Create(baseUri string) *Wttr {
	return &Wttr{baseUri}
}

func (w *Wttr) RetrieveWeather(city string, format string) (translator.Weather, error) {
	var body response
	query := url.Values{}
	query.Add("format", format)
	uri := fmt.Sprintf("%s/%s?%s", w.baseUri, city, query.Encode())

	resp, _ := http.DefaultClient.Get(uri)
	_ = json.NewDecoder(resp.Body).Decode(&body)
	tempInCelsius := body.CurrentCondition[0].TempC

	return translator.Weather{CurrentTempInCelsius: tempInCelsius}, nil
}

type response struct {
	CurrentCondition []struct {
		TempC string `json:"temp_C"`
	} `json:"current_condition"`
}
