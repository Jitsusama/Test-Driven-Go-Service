package translator_test

import (
	"fmt"
	"playing-around/pkg/translator"
	"reflect"
	"testing"
)

func TestTranslatesWttrRequests(t *testing.T) {
	for _, v := range []struct {
		city string
		temp string
	}{{"Milton", "24"}, {"Forest", "30"}} {
		t.Run(fmt.Sprintf("for city %s & temp %s", v.city, v.temp), func(t *testing.T) {
			// Arrange
			wttr := &mockWttr{stubbedResponse: translator.Weather{CurrentTempInCelsius: v.temp}}
			trans := translator.Create(wttr)

			// Act
			actualTemp, err := trans.RetrieveTemperature(v.city)

			// Assert
			if err != nil {
				t.Fatalf("temperature retrieval failed: %v", err)
			}
			if actualTemp != v.temp {
				t.Fatalf("expected temperature %v, got %v", v.temp, actualTemp)
			}
			if !reflect.DeepEqual(wttr.retrieveWeatherCalls, []retrieveWeather{{v.city, "j1"}}) {
				t.Fatalf(
					"expected retrieve weather called as %v, got %v",
					[]retrieveWeather{{v.city, "j1"}}, wttr.retrieveWeatherCalls,
				)
			}
		})
	}
}

type mockWttr struct {
	stubbedResponse      translator.Weather
	retrieveWeatherCalls []retrieveWeather
}

type retrieveWeather struct {
	city   string
	format string
}

func (w *mockWttr) RetrieveWeather(city string, format string) (translator.Weather, error) {
	w.retrieveWeatherCalls = append(w.retrieveWeatherCalls, retrieveWeather{city, format})
	return w.stubbedResponse, nil
}
