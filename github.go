package torvalds

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/google/go-github/v57/github"
)

type Event struct {
	Number     int        `json:"number"`
	Repository Repository `json:"repository"`
}
type Repository struct {
	Name  string `json:"name"`
	Owner Owner  `json:"owner"`
}
type Owner struct {
	Login string `json:"login"`
}

var ErrCouldNotFindEvent = errors.New("could not find pull request event")
var ErrCouldNotFindGithubToken = errors.New("could not find github token")

func GetDiffFromPullRequest() (string, error) {
	pathToEvent, ok := os.LookupEnv("GITHUB_EVENT_PATH")
	if !ok {
		return "", ErrCouldNotFindEvent
	}
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return "", ErrCouldNotFindEvent
	}

	data, _ := os.ReadFile(pathToEvent)
	var event Event
	json.Unmarshal(data, &event)

	client := github.NewClient(nil).WithAuthToken(token)
	bg := context.Background()
	diff, _, err := client.PullRequests.GetRaw(bg, event.Repository.Owner.Login, event.Repository.Name, event.Number, github.RawOptions{Type: github.Patch})
	if err != nil {
		return "", err
	}

	return diff, nil
}
