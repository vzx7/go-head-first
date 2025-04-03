package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", file, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %w", file, err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file1> [<file2> ...]", os.Args[0])
	}

	for _, file := range os.Args[1:] {
		if err := lineByLine(file); err != nil {
			log.Println(err)
		}

	}
}
