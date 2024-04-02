package v1Users

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	routesUsers := r.Group("/api/v1/users")

	routesUsers.POST("/", CreateUser)
	routesUsers.GET("/", GetUsers)
	routesUsers.GET("/:id", GetUser)
	routesUsers.PUT("/:id", UpdateUser)
	routesUsers.DELETE("/:id", DeleteUser)
}
