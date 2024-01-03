package main

import (
	"context"

	"github.com/google/go-github/v57/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

func main() {
	client := github.NewClient(nil).WithAuthToken("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")
	ctx := context.Background()

	prs, _, _ := client.PullRequests.List(ctx, "algleymi", "what-would-linus-torvalds-say", nil)

	print(*prs[0].Assignee.Login)
}
