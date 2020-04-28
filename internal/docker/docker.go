package docker

// import (
// 	"log"

// 	docker "github.com/docker/go-docker"
// 	"github.com/mactynow/trigger/internal/config"
// 	"golang.org/x/net/context"
// )

// type DockerEvent struct {
// 	Status         string `json:"status"`
// 	Error          string `json:"error"`
// 	Progress       string `json:"progress"`
// 	ProgressDetail struct {
// 		Current int `json:"current"`
// 		Total   int `json:"total"`
// 	} `json:"progressDetail"`
// }

// // PullImage pulls an image from a repository if it exists
// func PullImage(config *config.Config) {
// 	client, err := docker.NewEnvClient()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	events, err := client.ImagePull(context.Background(), config.ContainerName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
