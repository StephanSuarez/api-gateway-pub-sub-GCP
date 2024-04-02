package cmd

import (
	"fmt"
	"log"

	"github.com/StephanSuarez/chat-rooms/api-gateway/internal/common/conf"
	"github.com/StephanSuarez/chat-rooms/api-gateway/internal/v1Rooms"
	"github.com/StephanSuarez/chat-rooms/api-gateway/internal/v1Users"
	"github.com/gin-gonic/gin"
)

type App struct {
	Env    *conf.Env
	Router *gin.Engine
}

func NewApp() *App {

	app := &App{}

	app.Env = conf.NewEnv()

	app.Router = gin.Default()
	return app
}

func (app *App) Start() {
	addr := fmt.Sprintf("%s:%s", app.Env.IPAddress, app.Env.ServerAddress)
	log.Printf("Server is running on: %s", addr)

	v1Rooms.Router(app.Router)
	v1Users.Router(app.Router)

	err := app.Router.Run(app.Env.PortServer)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
