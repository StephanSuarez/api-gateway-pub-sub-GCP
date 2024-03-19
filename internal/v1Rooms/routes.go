package v1Rooms

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	routerRooms := r.Group("v1/rooms")

	routerRooms.GET("")
	routerRooms.POST("/", CreateUser)
	routerRooms.GET("/:id")
	routerRooms.PUT("/:id")
	routerRooms.DELETE("/:id")
}
