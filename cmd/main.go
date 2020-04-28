package main

import (
	"log"

	"github.com/mactynow/trigger/internal/config"
	"github.com/mactynow/trigger/internal/pubsub"
)

func main() {
	config := config.GetConfigFromEnv("trigger")

	log.Printf("Starting trigger on topic %s for project %s", config.TopicID, config.ProjectID)

	pubsub.StartReceiver(config)
}
