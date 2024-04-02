package v1Users

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	urlBase = os.Getenv("URL_USERS_BASE")
}

func CreateUser(ctx *gin.Context) {
	requestBody := bytes.Buffer{}
	_, err := io.Copy(&requestBody, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	microserviceURL := urlBase

	response, err := http.Post(microserviceURL, "application/json", &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al realizar la solicitud al microservicio",
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al leer la respuesta del microservicio",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func GetUsers(ctx *gin.Context) {
	microserviceURL := urlBase

	response, err := http.Get(microserviceURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al realizar la solicitud al microservicio",
		})
		return
	}
	defer response.Body.Close()

	var responseBody []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al leer la respuesta del microservicio",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	microserviceURL := fmt.Sprintf("%s/%s", urlBase, id)

	response, err := http.Get(microserviceURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al realizar la solicitud al microservicio",
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al leer la respuesta del microservicio",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func UpdateUser(ctx *gin.Context) {
	requestBody := bytes.Buffer{}
	_, err := io.Copy(&requestBody, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	id := ctx.Param("id")
	microserviceURL := fmt.Sprintf("%s/%s", urlBase, id)

	req, err := http.NewRequest("PUT", microserviceURL, &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud HTTP"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la solicitud al microservicio"})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la respuesta del microservicio"})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	microserviceURL := fmt.Sprintf("%s/%s", urlBase, id)

	request, err := http.NewRequest("DELETE", microserviceURL, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al crear la solicitud",
		})
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al realizar la solicitud al microservicio"})
		return
	}

	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al leer la respuesta del microservicio",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}
