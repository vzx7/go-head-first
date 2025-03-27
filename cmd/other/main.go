package main

import (
	"fmt"
	"sort"

	"github.com/vzx7/go-head-first/pkg/reflection"
	S "github.com/vzx7/go-head-first/pkg/sort"
)

var data = []S.S2{
	{F1: 1, F2: "One", F3: S.S1{F1: 1, F2: "S1_1", F3: 10}},
	{F1: 2, F2: "Two", F3: S.S1{F1: 1, F2: "S1_1", F3: 20}},
	{F1: -1, F2: "Two", F3: S.S1{F1: -1, F2: "S1_1", F3: -20}},
}

func main() {
	reflection.DoReflect()
	reflection.DoModifyByReflect()
	doSort(data)
}

func doSort(data []S.S2) {
	fmt.Println("before:", data)
	sort.Sort(S.S2slice(data))
	fmt.Println("after:", data)
	sort.Sort(sort.Reverse(S.S2slice(data)))
	fmt.Println("reverse:", data)
}
