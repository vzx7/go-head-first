package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)

	b = []byte("Byte slice $")
	fmt.Println("Byte slice", b)

	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	fmt.Println("Length of b:", len(b))

	fmt.Println("----------------------------")
	DoSliceStruct()
}

type record struct {
	Field1 int
	Field2 string
}

func DoSliceStruct() {
	S := []record{}

	for i := 0; i < 10; i++ {
		text := "text" + strconv.Itoa(i)
		temp := record{Field1: i, Field2: text}
		S = append(S, temp)
	}

	fmt.Println("Index 0:", S[0].Field1, S[0].Field2)
	fmt.Println("Number of structures:", len(S))
	sum := 0
	for _, k := range S {
		sum += k.Field1
	}
	fmt.Println("Sum:", sum)
}
