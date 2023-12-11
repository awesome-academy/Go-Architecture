package envutils

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "Go-Architecture"

// LoadEnv loads env vars from .env
func LoadEnv() error {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		return err
	}
	return nil
}
