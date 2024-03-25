package v1Rooms

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {

	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}
	if err := createUserPub(ctx.Writer, body); err != nil {
		log.Println(err)
	}
}

func DeleteUser(ctx *gin.Context) {
	roomID := ctx.Param("id")

	if roomID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	fmt.Println("sending")

	if err := deleteUserPub(ctx.Writer, roomID); err != nil {
		log.Println(err)
	}
}
