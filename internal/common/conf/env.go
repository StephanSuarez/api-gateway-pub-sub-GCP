package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	ServerAddress string
	PortServer    string
}

func NewEnv() *Env {
	env := Env{}

	if err := godotenv.Load(".env.yaml"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	env.PortServer = os.Getenv("PORT_SERVER")

	return &env
}
