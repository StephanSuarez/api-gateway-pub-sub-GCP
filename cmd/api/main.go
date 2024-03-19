package main

import (
	"log"

	"github.com/StephanSuarez/chat-rooms/api-gateway/cmd"
)

func main() {
	app := cmd.NewApp()

	log.Println(app)

	app.Start()
}
