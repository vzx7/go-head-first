package main

import (
	"fmt"
	"log"

	"github.com/vzx7/go-head-first/pkg/datafile"
)

func main() {
	lines, err := datafile.GetStrings("cmd/count/votes")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
