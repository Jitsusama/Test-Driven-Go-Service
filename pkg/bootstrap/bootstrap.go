package bootstrap

type App struct{}

func Create(port int, wttrUrl string) *App {
	return &App{}
}

func (a App) Start() {
}

func (a App) Stop() {
}
