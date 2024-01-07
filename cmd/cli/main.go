package main

import (
	"log"

	torvaldsify "github.com/arnoschutijzer/what-would-linus-torvalds-say"
)

func main() {
	event, err := torvaldsify.ParseEventFromGithubActionsEvent()
	diff, err := torvaldsify.GetDiffFromPullRequest(event.Repository.Owner.Login, event.Repository.Name, event.Number)
	if err != nil {
		log.Fatal(err)
	}

	whatTorvaldsSaid, err := torvaldsify.AskTorvalds(diff)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(whatTorvaldsSaid)
}
