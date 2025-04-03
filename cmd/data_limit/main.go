package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := io.ReadFull(f, buffer)

	if err == io.EOF || err == io.ErrUnexpectedEOF {
		return buffer[:n] // Возвращаем то, что успели прочитать
	}

	if err != nil {
		log.Printf("Error reading file: %v", err)
		return nil
	}

	return buffer
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s <buffer size> <filename>", os.Args[0])
	}

	bufferSize, err := strconv.Atoi(os.Args[1])
	if err != nil || bufferSize <= 0 {
		log.Fatalf("Invalid buffer size: %v", err)
	}

	file := os.Args[2]
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", file, err)
	}
	defer f.Close()

	readData := readSize(f, bufferSize)
	if len(readData) > 0 {
		fmt.Print(string(readData))
	}
	fmt.Println()
}
