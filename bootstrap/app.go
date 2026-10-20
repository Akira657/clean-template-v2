package bootstrap

import (
	"clean-template/db/postgres"
)

type Application struct {
	Env      *Env
	Postgres postgres.PostgresDatabase
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	InitSlog(app.Env)
	app.Postgres = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresConnection(app.Postgres)
}

func (app *Application) ClosePostgresDBConnection() {
	ClosePostgresConnection(app.Postgres)
	// if app.Postgres.DB != nil {
	// 	app.Postgres.DB.Close()
	// }
}
