package bootstrap

import (
	"playing-around/pkg/server"
	"playing-around/pkg/translator"
)

type App struct {
	svr *server.Server
}

func Create(port int, wttrUrl string) *App {
	trans := translator.Create(&dummyWeather{})
	svr := server.Create(port, trans)
	return &App{svr}
}

func (a App) Start() error {
	return a.svr.Start()
}

func (a App) Stop() error {
	return a.svr.Stop()
}

type dummyWeather struct {
}

func (t *dummyWeather) RetrieveWeather(city string, format string) (translator.Weather, error) {
	return translator.Weather{CurrentTempInCelsius: "wttr not implemented yet"}, nil
}
