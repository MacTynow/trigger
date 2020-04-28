package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/mactynow/trigger/internal/config"
)

// StartReceiver creates or resumes consuming messages from a subcription to a topic
func StartReceiver(config *config.Config) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatal(err.Error())
	}

	topic := client.Topic(config.TopicID)

	var sub = client.Subscription("trigger")
	ok, err := sub.Exists(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	if !ok {
		log.Print("Subscription doesn't exist, creating it...")
		sub, err = client.CreateSubscription(ctx, "trigger", pubsub.SubscriptionConfig{Topic: topic})
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		text := string(m.Data)
		log.Printf("Message received: %s", text)
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
