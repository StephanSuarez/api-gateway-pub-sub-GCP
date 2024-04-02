package conf

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	IPAddress     string
	ServerAddress string
	PortServer    string
	ProjectID     string
}

func NewEnv() *Env {
	env := Env{}

	if err := godotenv.Load(".env.yaml"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	env.PortServer = os.Getenv("PORT_SERVER")
	env.ProjectID = os.Getenv("PROJECT_ID")
	env.IPAddress = os.Getenv("IP_ADDRESS")

	return &env
}
