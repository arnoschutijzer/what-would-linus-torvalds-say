package main

import (
	"log"
	"os"
	"strconv"

	torvaldsify "github.com/arnoschutijzer/what-would-linus-torvalds-say"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: go run main.go owner repo prNumber")
	}
	owner := os.Args[1]
	repo := os.Args[2]
	pr := os.Args[3]

	prNumber, err := strconv.Atoi(pr)
	if err != nil {
		log.Fatal("prNumber is not a number")
	}

	diff, err := torvaldsify.GetDiffFromPullRequest(owner, repo, prNumber)
	if err != nil {
		log.Fatal(err)
	}

	whatTorvaldsSaid, err := torvaldsify.AskTorvalds(diff)
	if err != nil {
		log.Fatal(err)
	}

	err = torvaldsify.AddCommentToPullRequest(owner, repo, prNumber, whatTorvaldsSaid)
	if err != nil {
		log.Fatal(err)
	}
}
