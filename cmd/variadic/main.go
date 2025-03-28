package main

import (
	"fmt"
	"os"
)

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	sum := float64(0)
	for _, a := range s {
		sum += a
	}
	s[0] = 1000
	return sum
}

func everything(input ...any) {
	fmt.Println(input...)
}

func main() {
	sum := addFloats("Adding numbers [1]...", 1.3, 5.6, 76.0, 89.12, 899.98)
	fmt.Println("Sum", sum)
	s := []float64{12.12, 23.34, 3.9, 89.0}
	sum = addFloats("Adding numbers [2]...", s...)
	fmt.Println("Sum", sum)
	everything(s)

	empty := make([]any, len(os.Args[1:]))
	for i, v := range os.Args[1:] {
		empty[i] = v
	}
	everything(empty...)

	str := []string{"One", "Two", "Three"}
	everything(str, str, str)
}
