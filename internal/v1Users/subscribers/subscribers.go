package subscribers

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

var projectID string

func init() {
	if err := godotenv.Load(".env.yaml"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectID = os.Getenv("PROJECT_ID")
}

func CreateUserResponseSubs(w io.Writer) error {
	subID := "create-user-response-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}
	fmt.Fprintf(w, "Received %d messages\n", received)

	// return data

	return nil
}
