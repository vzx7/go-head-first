package main

import (
	"fmt"
	"log"

	"github.com/vzx7/go-head-first/pkg/datafile"
)

func main() {
	numbers, err := datafile.GoFloats("cmd/average/data")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
