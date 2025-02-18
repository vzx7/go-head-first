package main

import (
	"fmt"
	"log"

	"github.com/vzx7/go-head-first/pkg/keyboard"
)

func main() {
	fmt.Print("Enter a grabe: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
