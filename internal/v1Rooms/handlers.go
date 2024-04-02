package v1Rooms

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var urlBase string

func init() {
	if err := godotenv.Load(".env.yaml"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	urlBase = os.Getenv("URL_ROOMS_BASE")
}

func CreateRoom(ctx *gin.Context) {
	requestBody := bytes.Buffer{}
	_, err := io.Copy(&requestBody, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	microserviceURL := urlBase
	log.Println(microserviceURL)

	response, err := http.Post(microserviceURL, "application/json", &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al realizar la solicitud al microservicio",
		})
		return
	}
	defer response.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al leer la respuesta del microservicio",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody.String())
}
