package main

import (
	"log"

	gh "github.com/arnoschutijzer/what-would-linus-torvalds-say"
)

func main() {
	diff, err := gh.GetDiffFromPullRequest()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(*diff)
}
