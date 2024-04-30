package v1Rooms

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

	urlBase = os.Getenv("URL_ROOMS_BASE")
}

func CreateRoom(ctx *gin.Context) {
	requestBody := bytes.Buffer{}
	_, err := io.Copy(&requestBody, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error reading request body"})
		return
	}

	microserviceURL := urlBase

	response, err := http.Post(microserviceURL, "application/json", &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al realizar la solicitud al microservicio",
		})
		return
	}
	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error al leer la respuesta del microservicio",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func GetRooms(ctx *gin.Context) {
	microserviceURL := urlBase

	response, err := http.Get(microserviceURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error making request",
		})
		return
	}
	defer response.Body.Close()

	var responseBody []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading reponse",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func GetRoomByID(ctx *gin.Context) {
	roomID := ctx.Param("id")

	microserviceURL := fmt.Sprintf("%s/%s", urlBase, roomID)
	response, err := http.Get(microserviceURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error making request",
		})
		return
	}

	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading response",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func UpdateCompleteRoom(ctx *gin.Context) {
	requestBody := bytes.Buffer{}
	_, err := io.Copy(&requestBody, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading request body",
		})
		return
	}

	roomID := ctx.Param("id")
	microserviceURL := fmt.Sprintf("%s/%s", urlBase, roomID)

	log.Println(microserviceURL)
	req, err := http.NewRequest("PUT", microserviceURL, &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error making request",
		})
		return
	}

	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error sending request",
		})
		return
	}

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading response",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func UpdateParcialRoom(ctx *gin.Context) {
	requestBody := bytes.Buffer{}
	_, err := io.Copy(&requestBody, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading request body",
		})
		return
	}

	roomID := ctx.Param("id")
	microserviceURL := fmt.Sprintf("%s/%s", urlBase, roomID)

	log.Println(microserviceURL)
	req, err := http.NewRequest("PARCH", microserviceURL, &requestBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error making request",
		})
		return
	}

	req.Header.Set("Content-type", "application/json")
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error sending request",
		})
		return
	}

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading response",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}

func DeleteRoom(ctx *gin.Context) {
	roomID := ctx.Param("id")
	microserviceURL := fmt.Sprintf("%s/%s", urlBase, roomID)
	req, err := http.NewRequest("DELETE", microserviceURL, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error making request",
		})
		return
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error sending request",
		})
	}

	defer response.Body.Close()

	var responseBody map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error reading response",
		})
		return
	}

	ctx.JSON(response.StatusCode, responseBody)
}
