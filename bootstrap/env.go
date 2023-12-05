package bootstrap

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	AppEnv                 string
	ServerAddress          string
	DBHost                 string
	ContextTimeOut         int
	DBPort                 string
	DBUser                 string
	DBPass                 string
	DBName                 string
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
}

func NewEnv() *Env {
	env := Env{}

	env.AppEnv = os.Getenv("APP_ENV")
	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	contextTimeOut, err := strconv.ParseInt(os.Getenv("CONTEXT_TIMEOUT"), 10, 64)
	if err != nil {
		print(err)
	} else {
		env.ContextTimeOut = int(contextTimeOut)
	}
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("DB_NAME")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPass = os.Getenv("DB_PASSWORD")
	env.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	env.RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")
	accessTokenExpiryHour, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"), 10, 64)
	if err != nil {
		log.Println(err)
	} else {
		env.AccessTokenExpiryHour = int(accessTokenExpiryHour)
	}
	refreshTokenExpiryHour, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"), 10, 64)
	if err != nil {
		log.Println(err)
	} else {
		env.RefreshTokenExpiryHour = int(refreshTokenExpiryHour)
	}

	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &env
}
