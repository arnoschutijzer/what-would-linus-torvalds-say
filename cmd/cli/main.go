package main

import (
	"log"

	torvaldsify "github.com/arnoschutijzer/what-would-linus-torvalds-say"
)

func main() {
	diff, err := torvaldsify.GetDiffFromPullRequest()
	if err != nil {
		log.Fatal(err)
	}

	whatTorvaldsSaid, err := torvaldsify.AskTorvalds(diff)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(*whatTorvaldsSaid)
}
