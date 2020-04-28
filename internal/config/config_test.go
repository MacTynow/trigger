package config

import (
	"math/rand"
	"os"
	"testing"
)

const (
	TestProjectID     = string("test_projectid")
	TestTopicID       = string("test_topicid")
	TestContainerName = string("test_containername")
)

func TestGetConfigFromEnv(t *testing.T) {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	testAppName := string(b)

	os.Setenv(testAppName+"_PROJECTID", TestProjectID)
	os.Setenv(testAppName+"_TOPICID", TestTopicID)
	os.Setenv(testAppName+"_CONTAINERNAME", TestContainerName)

	testConfig := GetConfigFromEnv(testAppName)

	if testConfig.ProjectID != TestProjectID {
		t.Errorf("ProjectID, got: %s, want: %s.", testConfig.ProjectID, TestProjectID)
	} else if testConfig.TopicID != TestTopicID {
		t.Errorf("TopicID, got: %s, want: %s.", testConfig.TopicID, TestProjectID)
	} else if testConfig.ContainerName != TestContainerName {
		t.Errorf("ContainerName, got: %s, want: %s.", testConfig.ContainerName, TestContainerName)
	}
}
