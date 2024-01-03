package main

import (
	"context"

	"github.com/google/go-github/v57/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func main() {
	client := github.NewClient(nil).WithAuthToken("a_token")
	ctx := context.Background()

	prs, _, _ := client.PullRequests.List(ctx, "algleymi", "what-would-linus-torvalds-say", nil)

	print(*prs[0].Assignee.Login)
}
