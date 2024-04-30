package v1Rooms

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	routerRooms := r.Group("api/v1/rooms")

	routerRooms.GET("", GetRooms)
	routerRooms.POST("", CreateRoom)
	routerRooms.GET(":id", GetRoomByID)
	routerRooms.PUT(":id", UpdateCompleteRoom)
	routerRooms.PATCH(":id", UpdateParcialRoom)
	routerRooms.DELETE(":id", DeleteRoom)
}
