package main

import (
	"log"

	torvaldsify "github.com/arnoschutijzer/what-would-linus-torvalds-say"
)

func main() {
	event, err := torvaldsify.ParseEventFromGithubActionsEvent()
	if err != nil {
		log.Fatal(err)
	}

	diff, err := torvaldsify.GetDiffFromPullRequest(event.Repository.Owner.Login, event.Repository.Name, event.Number)
	if err != nil {
		log.Fatal(err)
	}

	whatTorvaldsSaid, err := torvaldsify.AskTorvalds(diff)
	if err != nil {
		log.Fatal(err)
	}

	err = torvaldsify.AddCommentToPullRequest(event.Repository.Owner.Login, event.Repository.Name, event.Number, whatTorvaldsSaid)
	if err != nil {
		log.Fatal(err)
	}
}
