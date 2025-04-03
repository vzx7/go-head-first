package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	if err != nil {
		log.Fatalf("Failed to open /dev/random: %v", err)
	}
	defer f.Close()

	var seed int64
	if err := binary.Read(f, binary.LittleEndian, &seed); err != nil {
		log.Fatalf("Failed to read seed: %v", err)
	}

	fmt.Println("Seed:", seed)
}
