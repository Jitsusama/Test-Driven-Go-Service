package bootstrap

import "playing-around/pkg/server"

type App struct {
	svr *server.Server
}

func Create(port int, wttrUrl string) *App {
	svr := server.Create(port)
	return &App{svr}
}

func (a App) Start() error {
	return a.svr.Start()
}

func (a App) Stop() error {
	return a.svr.Stop()
}
