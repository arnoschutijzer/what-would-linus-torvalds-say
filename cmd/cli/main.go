package main

import (
	"context"
	"os"

	"github.com/google/go-github/v57/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func main() {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		os.Exit(1)
	}

	client := github.NewClient(nil).WithAuthToken(token)
	ctx := context.Background()

	prs, _, _ := client.PullRequests.List(ctx, "algleymi", "what-would-linus-torvalds-say", &github.PullRequestListOptions{State: "all"})

	print(*prs[0].Title)
}
