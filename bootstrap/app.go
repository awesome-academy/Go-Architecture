package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env *Env
	Db  *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDB(app.Db)
}
