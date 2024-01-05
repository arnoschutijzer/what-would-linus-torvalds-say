package main

import (
	"log"
	"os"

	gh "github.com/arnoschutijzer/what-would-linus-torvalds-say"
)

func main() {
	diff, err := gh.GetDiffFromPullRequest()
	if err != nil {
		log.Fatal(err)
	}

	if len(*diff) > 10_000 {
		log.Printf("diff might be too large since diff length (%d) > 15_000 characters\n", len(*diff))
		log.Println("failing silently...")
		os.Exit(0)
	}

	log.Println(diff)
}
