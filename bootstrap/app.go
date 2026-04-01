package bootstrap

import (
	"database/sql"
)

type Application struct {
	Env      *Env
	Postgres *sql.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}
