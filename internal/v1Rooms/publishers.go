package v1Rooms

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/pubsub"
)

func createUserPub(w io.Writer, body []byte) error {
	projectID := "monograph-417222"
	topicID := "create-user"

	log.Println("into pub")
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
