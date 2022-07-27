package bootstrap

import (
	"playing-around/pkg/server"
	"playing-around/pkg/translator"
	"playing-around/pkg/wttr"
)

type App struct {
	svr *server.Server
}

func Create(port int, wttrUrl string) *App {
	w := wttr.Create(wttrUrl)
	trans := translator.Create(w)
	svr := server.Create(port, trans)
	return &App{svr}
}

func (a App) Start() error {
	return a.svr.Start()
}

func (a App) Stop() error {
	return a.svr.Stop()
}
