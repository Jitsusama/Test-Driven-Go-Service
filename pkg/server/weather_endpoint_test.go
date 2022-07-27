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
	port := freeTcpPort()
	city := "Milton"
	temperature := "24"

	// Arrange
	trans := &mockedTranslator{stubbedResponse: temperature}

	svr := server.Create(port, trans)
	go func() { _ = svr.Start() }()
	defer func() { _ = svr.Stop() }()

	// Act
	req, _ := http.NewRequest(
		"GET", fmt.Sprintf("http://localhost:%d/weather/%s", port, city), nil,
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
		t.Fatalf("expected content type %s, not %s", "text/plain", resp.Header.Get("Content-Type"))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(body) != temperature {
		t.Fatalf("expected body %s to be %s", string(body), temperature)
	}
	if !reflect.DeepEqual(trans.retrieveWeathers, []string{city}) {
		t.Fatalf("expected 1 call with city %s, got %v", city, trans.retrieveWeathers)
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
