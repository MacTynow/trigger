package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
)

// Config from environment definition
type Config struct {
	ProjectID string
	TopicID   string
}

func main() {
	var config Config
	err := envconfig.Process("trigger", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Starting trigger on topic %s for project %s", config.TopicID, config.ProjectID)

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
