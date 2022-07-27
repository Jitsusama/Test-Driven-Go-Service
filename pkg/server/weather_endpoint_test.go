package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"playing-around/pkg/server"
	"reflect"
	"testing"
)

func TestPassesRequestToTranslationLayer(t *testing.T) {
	for _, v := range []struct {
		port int
		city string
		temp string
	}{
		{port: freeTcpPort(), city: "Milton", temp: "24"},
		{port: freeTcpPort(), city: "Forest", temp: "30"},
	} {
		t.Run(fmt.Sprintf("for city %s and temp %s", v.city, v.temp), func(t *testing.T) {
			// Arrange
			trans := &mockedTranslator{stubbedResponse: v.temp}

			svr := server.Create(v.port, trans)
			go func() { _ = svr.Start() }()
			defer func() { _ = svr.Stop() }()

			// Act
			req, _ := http.NewRequest(
				"GET", fmt.Sprintf("http://localhost:%d/weather/%s", v.port, v.city), nil,
			)
			resp, err := http.DefaultClient.Do(req)

			// Assert
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("expected status code %d, not %d", http.StatusOK, resp.StatusCode)
			}
			if resp.Header.Get("Content-Type") != "text/plain" {
				t.Fatalf(
					"expected content type %s, not %s", "text/plain", resp.Header.Get("Content-Type"),
				)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil || string(body) != v.temp {
				t.Fatalf("expected body %s to be %s", string(body), v.temp)
			}
			if !reflect.DeepEqual(trans.retrieveWeathers, []string{v.city}) {
				t.Fatalf("expected 1 call with city %s, got %v", v.city, trans.retrieveWeathers)
			}
		})
	}
}

type mockedTranslator struct {
	stubbedResponse  string
	retrieveWeathers []string
}

func (t *mockedTranslator) RetrieveWeather(city string) string {
	t.retrieveWeathers = append(t.retrieveWeathers, city)
	return t.stubbedResponse
}
