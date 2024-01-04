package main

import (
	"encoding/json"
	"log"
	"os"
	// with go modules enabled (GO111MODULE=on or outside GOPATH)
)

type Event struct {
	PullRequest PullRequest `json:"pull_request"`
}
type PullRequest struct {
	DiffURL string `json:"diff_url"`
}

func main() {
	pathToEvent, ok := os.LookupEnv("GITHUB_EVENT_PATH")
	if !ok {
		log.Fatal("could not find event, can't do anything")
	}

	data, _ := os.ReadFile(pathToEvent)
	var event Event
	json.Unmarshal(data, &event)
}
