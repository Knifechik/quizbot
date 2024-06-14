package app

type App struct {
	repo Repo
}

func New(r Repo) *App {
	return &App{
		repo: r,
	}
}
