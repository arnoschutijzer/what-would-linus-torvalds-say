package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/google/go-github/v57/github"
)

type Event struct {
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}
type PullRequest struct {
	DiffURL string `json:"diff_url"`
}
type Repository struct {
	Name  string `json:"name"`
	Owner Owner  `json:"owner"`
}
type Owner struct {
	Login string `json:"login"`
}

func main() {
	pathToEvent, ok := os.LookupEnv("GITHUB_EVENT_PATH")
	if !ok {
		log.Fatal("could not find event, can't do anything")
	}
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		log.Fatal("could not find token, can't do anything")
	}

	data, _ := os.ReadFile(pathToEvent)
	var event Event
	json.Unmarshal(data, &event)

	client := github.NewClient(nil).WithAuthToken(token)
	bg := context.Background()
	diff, _, err := client.PullRequests.GetRaw(bg, event.Repository.Owner.Login, event.Repository.Name, event.Number, github.RawOptions{Type: github.Diff})
	if err != nil {
		log.Fatal("could not get diff due to: ", err)
	}

	if len(diff) > 10_000 {
		log.Printf("yikes, diff might be too large %d > 10_000\n", len(diff))
		log.Println("failing silently...")
		os.Exit(0)
	}

	log.Println(diff)
}
