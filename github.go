package torvalds

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/google/go-github/v61/github"
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

func ParseEventFromGithubActionsEvent() (Event, error) {
	pathToEvent, ok := os.LookupEnv("GITHUB_EVENT_PATH")
	if !ok {
		return Event{}, ErrCouldNotFindEvent
	}

	return parseEventFromFile(pathToEvent)
}

func parseEventFromFile(path string) (Event, error) {
	data, _ := os.ReadFile(path)

	var event Event
	err := json.Unmarshal(data, &event)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func GetDiffFromPullRequest(owner string, repositoryName string, prNumber int) (string, error) {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return "", ErrCouldNotFindGithubToken
	}

	client := github.NewClient(nil).WithAuthToken(token)
	bg := context.Background()
	diff, _, err := client.PullRequests.GetRaw(bg, owner, repositoryName, prNumber, github.RawOptions{Type: github.Patch})
	if err != nil {
		return "", err
	}

	return diff, nil
}

func AddCommentToPullRequest(owner string, repositoryName string, prNumber int, comment string) error {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return ErrCouldNotFindGithubToken
	}

	client := github.NewClient(nil).WithAuthToken(token)
	bg := context.Background()
	_, _, err := client.Issues.CreateComment(bg, owner, repositoryName, prNumber, &github.IssueComment{Body: &comment})
	if err != nil {
		return err
	}

	return nil
}
