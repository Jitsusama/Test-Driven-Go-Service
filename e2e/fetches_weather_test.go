package e2e

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"playing-around/pkg/bootstrap"
	"reflect"
	"testing"
	"time"
)

func TestGetsCurrentWeatherInCelsius(t *testing.T) {
	var wttrRequests []wttrRequest
	listeningPort := freeTcpPort()

	// Arrange
	wttr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	app := bootstrap.Create(listeningPort, wttr.URL)
	go func() { _ = app.Start() }()
	defer func() { _ = app.Stop() }()
	time.Sleep(time.Millisecond)

	// Act
	req, _ := http.NewRequest(
		"GET", fmt.Sprintf("http://localhost:%d/weather?for=Milton", listeningPort), nil,
	)
	resp, err := http.DefaultClient.Do(req)

	// Assert
	if err != nil {
		t.Fatalf("http request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status code to be %d, not %d", http.StatusOK, resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "text/plain" {
		t.Fatalf(
			"expected content type of %v, but got %v",
			"text/plain", resp.Header.Get("Content-Type"),
		)
	}

	actualTemp, _ := ioutil.ReadAll(resp.Body)
	if string(actualTemp) != "24" {
		t.Fatalf("actualTemp should be %v but was %v", "24", string(actualTemp))
	}

	expected := wttrRequest{method: "GET", path: "/Milton", responseFormat: "j1"}
	if len(wttrRequests) != 1 || !reflect.DeepEqual(wttrRequests[0], expected) {
		t.Fatalf("expected 1 request matching %v, but got %v", expected, wttrRequests)
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

func freeTcpPort() int {
	l, _ := net.Listen("tcp4", ":0")
	defer func(l net.Listener) { _ = l.Close() }(l)
	return l.Addr().(*net.TCPAddr).Port
}
