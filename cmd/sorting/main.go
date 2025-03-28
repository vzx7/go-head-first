package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grades{
		{"J.", "Lewis", 10},
		{"M.", "Jons", 23},
		{"Y.", "Authron", 11},
	}

	less := func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	}

	isSorted(data, less)
	sort.Slice(data, less)
	isSorted(data, less)
}

func isSorted(data []Grades, less func(i, j int) bool) {
	isSorted := sort.SliceIsSorted(data, less)

	if isSorted {
		fmt.Println("It is sorted...")
	} else {
		fmt.Println("It is NOT sorted...")
	}
	fmt.Println("By Grade:", data)
}
