package main

import (
	"context"

	"github.com/google/go-github/v57/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func main() {
	client := github.NewClient(nil).WithAuthToken("a_token")
	ctx := context.Background()

	prs, _, _ := client.PullRequests.List(ctx, "algleymi", "what-would-linus-torvalds-say", &github.PullRequestListOptions{State: "all"})

	print(*prs[0].Title)
}
