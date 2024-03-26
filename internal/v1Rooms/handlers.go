package v1Rooms

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

// func CreateUser(ctx *gin.Context) {

// 	body, err := ctx.GetRawData()
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
// 		return
// 	}
// 	if err := createUserPub(ctx.Writer, body); err != nil {
// 		log.Println(err)
// 	}
// }

type Message struct {
	Content string `json:"content"`
}

func CreateUser(ctx *gin.Context) {
	conn, err := amqp.Dial("amqps://uaixaaut:4Hy7WEJGlEDZUEvDHXL91fSXEiv5gmU9@shrimp.rmq.cloudamqp.com/uaixaaut")
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error al abrir el canal: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"microservice_queue", // Nombre de la cola
		false,                // Durabilidad
		false,                // Eliminar cuando no hay consumidores
		false,                // Exclusiva
		false,                // No espera en colas
		nil,                  // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %v", err)
	}

	// Mensaje a enviar al microservicio
	msg := Message{Content: "Hola desde la API Gateway"}
	body, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error al serializar el mensaje: %v", err)
	}

	err = ch.Publish(
		"",     // Intercambio
		q.Name, // Ruta de la cola
		false,  // Mandar asincrónicamente
		false,  // Publicar al servidor si no hay consumidores
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Error al publicar el mensaje: %v", err)
	}

	log.Println("Mensaje enviado al microservicio")

	response := gin.H{"response": "No se recibió respuesta del microservicio"}

	// Canal para recibir la respuesta del microservicio
	responseChan := make(chan string)

	// Ejecutar la función WaitForResponse() en una goroutine
	go func() {
		responseFromMicroservice := WaitForResponse()
		responseChan <- responseFromMicroservice
	}()

	// Esperar hasta 5 segundos para recibir la respuesta del microservicio
	select {
	case <-time.After(5 * time.Second):
		// Si no se recibe respuesta dentro de los 5 segundos, no hagas nada y continúa con la respuesta predeterminada
	case responseFromMicroservice := <-responseChan:
		// Si se recibe una respuesta del microservicio, actualiza la respuesta
		if responseFromMicroservice != "" {
			response = gin.H{"response": responseFromMicroservice}
		}
	}

	// Continuar con el manejo de la solicitud del cliente
	ctx.JSON(http.StatusOK, response)
}

func WaitForResponse() string {
	responseChan := make(chan string)

	go func() {
		conn, err := amqp.Dial("amqps://uaixaaut:4Hy7WEJGlEDZUEvDHXL91fSXEiv5gmU9@shrimp.rmq.cloudamqp.com/uaixaaut")
		if err != nil {
			log.Fatalf("Error al conectar a RabbitMQ: %v", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Error al abrir el canal: %v", err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"response_queue", // Nombre de la cola de respuesta (debe coincidir con el nombre en el microservicio)
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Error al declarar la cola de respuesta: %v", err)
		}

		msgs, err := ch.Consume(
			q.Name,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Error al consumir mensajes de respuesta: %v", err)
		}

		// Esperar por la respuesta
		for d := range msgs {
			var response Message
			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				log.Fatalf("Error al decodificar la respuesta: %v", err)
			}
			log.Printf("Recibida respuesta del microservicio: %s", response.Content)

			// Enviar el contenido del mensaje recibido a través del canal
			responseChan <- response.Content
		}
	}()

	// Esperar activamente la respuesta del microservicio
	response := <-responseChan
	return response
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
