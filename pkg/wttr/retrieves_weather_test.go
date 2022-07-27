package wttr_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"playing-around/pkg/translator"
	"playing-around/pkg/wttr"
	"reflect"
	"testing"
)

func TestRetrievesWeatherFromWttr(t *testing.T) {
	var wttrRequests []wttrRequest
	temp := "24"
	city := "Milton"
	format := "j1"

	// Arrange
	wttrApi := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wttrRequests = append(wttrRequests, wttrRequest{
			method:         r.Method,
			path:           r.URL.Path,
			responseFormat: r.URL.Query().Get("format"),
		})
		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(wttrResponse{
			CurrentCondition: []wttrCondition{{TempC: "24"}},
		})
		_, _ = w.Write(body)
	}))

	w := wttr.Create(wttrApi.URL)

	// Act
	actualWeather, err := w.RetrieveWeather(city, format)

	// Assert
	if err != nil {
		t.Fatalf("failed to retrieve weather: %v", err)
	}
	expectedWeather := translator.Weather{CurrentTempInCelsius: temp}
	if !reflect.DeepEqual(actualWeather, expectedWeather) {
		t.Fatalf("expected weather %v, got %v", expectedWeather, actualWeather)
	}
	expectedRequests := []wttrRequest{{method: "GET", path: "/" + city, responseFormat: format}}
	if !reflect.DeepEqual(wttrRequests, expectedRequests) {
		t.Fatalf("expected %v request, but got %v", expectedRequests, wttrRequests)
	}
}

type wttrRequest struct {
	method         string
	path           string
	responseFormat string
}

type wttrResponse struct {
	CurrentCondition []wttrCondition `json:"current_condition"`
}

type wttrCondition struct {
	TempC string `json:"temp_C"`
}
