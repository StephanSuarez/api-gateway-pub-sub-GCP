package v1Users

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	routesUsers := r.Group("/v1/users")

	routesUsers.POST("/")
	routesUsers.GET("/")
	routesUsers.GET("/:id")
	routesUsers.PUT("/:id")
	routesUsers.DELETE("/:id")
	routesUsers.GET("/74abc")
}
