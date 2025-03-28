package main

import "fmt"

func fd1() {
	for i := 3; i >= 0; i-- {
		defer fmt.Print(" ", i)
	}
}

func fd2() {
	for i := 3; i >= 0; i-- {
		defer func() {
			fmt.Print(" ", i)
		}()
	}
}

func fd3() {
	for i := 3; i >= 0; i-- {
		defer func(n int) {
			fmt.Print(" ", n)
		}(i)
	}
}

func main() {
	fmt.Println("fd1 call:")
	fd1()
	fmt.Println("\nfd2 call:")
	fd2()
	fmt.Println("\nfd3 call:")
	fd3()
	fmt.Println("")
}
