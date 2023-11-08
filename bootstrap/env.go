package bootstrap

import (
	"Go-Architecture/utils/env"
	envparse "github.com/caarlos0/env/v10"
	"log"
)

type Env struct {
	AppEnv        string `env:"APP_ENV,required"`
	ServerAddress string `env:"SERVER_ADDRESS,required"`
	DBHost        string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort        string `env:"DB_PORT"`
	DBUser        string `env:"DB_USER"`
	DBPass        string `env:"DB_PASS"`
	DBName        string `env:"DB_NAME"`
}

func NewEnv() *Env {
	err := envutils.LoadEnv()
	if err != nil {
		log.Fatalf("Unable to load .env file: %e", err)
	}

	env := Env{}
	err = envparse.Parse(&env) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
