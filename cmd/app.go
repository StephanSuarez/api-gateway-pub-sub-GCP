package cmd

import (
	"fmt"
	"log"

	"github.com/StephanSuarez/chat-rooms/api-gateway/internal/common/conf"
	"github.com/StephanSuarez/chat-rooms/api-gateway/internal/v1Users"

	// "github.com/StephanSuarez/chat-rooms/api-gateway/internal/v1Rooms"
	"github.com/gin-contrib/cors"

	// "github.com/gin-contrib/cors" // Importa el middleware de CORS
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

	addr := fmt.Sprintf("%s:%s", app.Env.IPAddress, app.Env.ServerAddress)
	log.Printf("Server is running on: %s", addr)

	// app.Router.Use(middleware.CorsMiddleware())
	app.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	v1Users.Router(app.Router)

	err := app.Router.Run(app.Env.PortServer)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	return app
}

// func (app *App) Start() {
// 	addr := fmt.Sprintf("%s:%s", app.Env.IPAddress, app.Env.ServerAddress)
// 	log.Printf("Server is running on: %s", addr)

// 	// app.Router.Use(middleware.CorsMiddleware())
// 	app.Router.Use(cors.Default())

// 	// v1Rooms.Router(app.Router)
// 	routesUsers := app.Router.Group("/api/v1/users")

// 	routesUsers.POST("/", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"msg": "succ"})
// 	})

// 	// routesUsers.POST("/", v1Users.CreateUser)
// 	routesUsers.GET("/", v1Users.GetUsers)
// 	routesUsers.GET("/:id", v1Users.GetUser)
// 	routesUsers.PUT("/:id", v1Users.UpdateUser)
// 	routesUsers.DELETE("/:id", v1Users.DeleteUser)

// 	err := app.Router.Run(app.Env.PortServer)
// 	if err != nil {
// 		log.Fatalf("Error al iniciar el servidor: %v", err)
// 	}
// }
