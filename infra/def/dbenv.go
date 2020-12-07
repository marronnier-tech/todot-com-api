package def

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	DBAdmin   = "mysql"
	Protocol  = "tcp(127.0.0.1:3306)"
	Database  = "daily_todo"
	Charset   = "utf8mb4"
	ParseTime = "True"
	Loc       = "Local"
)

type FromEnv struct {
	User string
	Pass string
}

func GetEnv() (out FromEnv, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}
	out = FromEnv{
		User: os.Getenv("USER"),
		Pass: os.Getenv("ENVPASS"),
	}
	return
}
