package v1Rooms

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	// "github.com/StephanSuarez/chat-rooms/api-gateway/internal/v1Rooms/subscribers"
	"github.com/joho/godotenv"
)

var projectID string

func init() {
	if err := godotenv.Load(".env.yaml"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectID = os.Getenv("PROJECT_ID")
}

func createUserPub(w io.Writer, body []byte) error {
	topicID := "create-user"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: body,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}

	fmt.Fprintf(w, "Published message; msg ID: %v\n", id)

	return nil
}

func deleteUserPub(w io.Writer, roomID string) error {
	topicID := "delete-user"

	fmt.Println(roomID)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)

	messageData := fmt.Sprintf(`{"RoomID": "%s"}`, roomID)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(messageData),
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}

	fmt.Fprintf(w, "Published message; msg ID: %v\n", id)
	// subscribers.CreateUserSubs(os.Stdout)

	return nil
}
